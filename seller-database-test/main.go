package main

import (
	"context"
	"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	L "github.com/themaxermister/extras/seller-database-test/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	c := L.GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	var results []*L.Seller
	page := int64(1)
	limit := int64(5)
	target := "username"
	value := "ha"
	results = L.PageSellers(c, page, limit, bson.M{"effectivedate": -1}, bson.M{"default": true, fmt.Sprintf("%s", target): primitive.Regex{Pattern: fmt.Sprintf(`%s`, value), Options: "i"}})

	for _, item := range results {
		fmt.Println(item.Username)
	}
}

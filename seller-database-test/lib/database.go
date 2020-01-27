package lib

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Make func reusable

func GetClient() *mongo.Client {
	connectString := fmt.Sprint("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(connectString)
	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Client Error: ", err)
	}
	err = c.Connect(context.Background())
	if err != nil {
		log.Fatal("Connect Error: ", err)
	}
	return c
}

// Global Functions
func GetDistinctList(collect string, client *mongo.Client, targetCol string, filter bson.M) []interface{} {
	collection := client.Database("sellerdata").Collection(collect)
	result, err := collection.Distinct(context.TODO(), targetCol, filter)
	if err != nil {
		fmt.Println("Address Error: ", err)
	}
	return result
}

func DeleteDocument(collect string, client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("sellerdata").Collection(collect)
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting item: ", err)
	}
	return deleteResult.DeletedCount
}

func UpdateDocument(collect string, c *mongo.Client, updatedData interface{}, filter bson.M) int64 {
	collection := c.Database("sellerdata").Collection(collect)
	atualizacao := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating item: ", err)
	}
	return updatedResult.ModifiedCount
}

func AddOneArrayItem(collect string, client *mongo.Client, updatedData interface{}, filter bson.M) int64 {
	collection := client.Database("sellerdata").Collection(collect)
	atualizacao := bson.D{{Key: "$push", Value: updatedData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on adding one Hero: ", err)
	}
	return updatedResult.ModifiedCount
}

func CountTotal(collect string, c *mongo.Client, filter bson.M) int64 {
	collection := c.Database("sellerdata").Collection(collect)
	itemCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error on Counting all documents: ", err)
	}

	return itemCount
}

func RemoveOneArrayItem(collect string, client *mongo.Client, updatedData interface{}, filter bson.M) int64 {
	collection := client.Database("sellerdata").Collection(collect)
	atualizacao := bson.D{{Key: "$pull", Value: updatedData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on adding one Hero: ", err)
	}
	return updatedResult.ModifiedCount
}

// Seller Functions
func PageSellers(c *mongo.Client, pageNo int64, perPage int64, sort bson.M, filter bson.M) []*Seller {
	var sellers []*Seller
	findOptions := options.Find()
	findOptions.SetLimit(perPage)
	skip := (pageNo - 1) * perPage
	findOptions.SetSkip(skip)
	findOptions.SetSort(sort)
	collection := c.Database("sellerdata").Collection("sellerinfo")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal("Error on Finding all the documents: ", err)
	}
	for cur.Next(context.TODO()) {
		var seller Seller
		err = cur.Decode(&seller)
		if err != nil {
			log.Fatal("Error on Decoding the document: ", err)
		}
		sellers = append(sellers, &seller)
	}
	return sellers
}

func ReturnAllSellers(c *mongo.Client, filter bson.M) []*Seller {
	var sellers []*Seller
	collection := c.Database("sellerdata").Collection("sellerinfo")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error on Finding all Seller documents: ", err)
	}
	for cur.Next(context.TODO()) {
		var seller Seller
		err = cur.Decode(&seller)
		if err != nil {
			fmt.Println("Error on Decoding the document: ", err)
		}
		sellers = append(sellers, &seller)
	}
	return sellers
}

func ReturnOneSeller(c *mongo.Client, filter bson.M) Seller {
	var seller Seller
	collection := c.Database("sellerdata").Collection("sellerinfo")
	err := collection.FindOne(context.TODO(), filter).Decode(&seller)
	if err != nil {
		fmt.Println("Seller Return Error: ", err)
	}
	return seller
}

func InsertNewSeller(c *mongo.Client, seller Seller) interface{} {
	collection := c.Database("sellerdata").Collection("sellerinfo")
	insertResult, err := collection.InsertOne(context.TODO(), seller)
	if err != nil {
		log.Fatalln("Error on inserting new Seller: ", err)
	}
	return insertResult.InsertedID
}

// ADDRESS INFO
func ReturnOneAddress(c *mongo.Client, filter bson.M) AddressDoc {
	var file AddressDoc
	collection := c.Database("sellerdata").Collection("addressinfo")
	err := collection.FindOne(context.TODO(), filter).Decode(&file)
	if err != nil {
		fmt.Println("File Return Error: ", err)
	}
	return file
}

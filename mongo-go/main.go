package main

import (
	"encoding/json"
	"context"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Hero struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Type	string	`json:"type"`
	Signed bool   `json:"signed"`
	Villians []string `json:"villians"`
}

type Company struct {
	CompanyName string `json:"company`
	Heroes  []Hero `json:"heroes"`
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Company {
	var heroes []*Company
	collection := client.Database("civilact").Collection("heros")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero Company
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

func ReturnOneHero(client *mongo.Client, filter bson.M) Company {
	var hero Company
	collection := client.Database("civilact").Collection("heros")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func InsertNewHero(client *mongo.Client, hero Company) interface{} {
	collection := client.Database("civilact").Collection("heros")
	insertResult, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Company", err)
	}
	return insertResult.InsertedID
}

func RemoveOneCompany(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heros")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Company", err)
	}
	return deleteResult.DeletedCount
}

func UpdateHero(client *mongo.Client, updatedData interface{}, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heros")
	atualizacao := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating one Company: ", err)
	}
	return updatedResult.ModifiedCount
}

func MongoToString(seller interface{}) string {
	json, err := json.Marshal(seller)

	if err != nil {
		fmt.Printf("JSON Error: %v on %s", err, seller)
		return ""
	}

	return string(json)
}

func RemoveOneArrayItem(client *mongo.Client, updatedData interface{}, filter bson.M) int64{
	collection := client.Database("civilact").Collection("heros")
	atualizacao := bson.D{{Key: "$pull", Value: updatedData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on adding one Hero: ", err)
	}
	return updatedResult.ModifiedCount
}

func main() {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	
	/*heroes := ReturnAllHeroes(c, bson.M{"heroes.villians" : bson.A{"Thanos", "T-Ray"}})
	//heroes := ReturnAllHeroes(c, bson.M{"signed": true})
	for _, hero := range heroes {
		for _, item := range hero.Heroes{
			fmt.Println(item.Name)
		}
	}

	
	hero := ReturnOneHero(c, bson.M{"name": "Vision"})
	log.Println(hero.Name, hero.Alias, hero.Signed)


	var hero = Company{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: true}
	insertedID := InsertNewHero(c, hero)
	log.Println(insertedID)
	hero = ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	log.Println(hero.Name, hero.Alias, hero.Signed)


	heroesRemoved := RemoveOneHero(c, bson.M{"Alias": "Doctor Strange"})
	log.Println("Heroes removed count:", heroesRemoved)
	var hero = ReturnOneHero(c, bson.M{"Alias": "Doctor Strange"})
	log.Println("Is Company empty?", hero == Company{})*/
	
	var newHero = Hero{Name: "Clark Kent", Alias: "Superman", Signed: true, Villians:[]string{"Item1, Item3"}}
	heroesUpdated := UpdateHero(c, bson.M{fmt.Sprintf("heroes.%s", "$"): newHero}, bson.M{"heroes.alias" : "Doctor Strange"})
	log.Println("Heroes updated count:", heroesUpdated)
	//hero := ReturnOneHero(c, bson.M{"heroes.alias": "Doctor Strange"})

	//var hero = Hero{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: false,  Villians:[]string{"Item1", "Item2"}}
	//AddOneArrayItem(c, bson.M{"heroes": hero}, bson.M{"company" : "Marvel"})
	//test := RemoveOneArrayItem(c, bson.M{"heroes": bson.M{"name": "Stephen Strange"}}, bson.M{"company": "Marvel"})
	//log.Println("Heroes updated count:", test)
	
	company := ReturnOneHero(c, bson.M{"company": "Marvel"})
	for _, item := range company.Heroes {
		fmt.Println(item.Type)
	}
}

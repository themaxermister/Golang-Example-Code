package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Hero struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
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

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Hero {
	var heroes []*Hero
	collection := client.Database("civilact").Collection("heros")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero Hero
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

func ReturnOneHero(client *mongo.Client, filter bson.M) Hero {
	var hero Hero
	collection := client.Database("civilact").Collection("heros")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func InsertNewHero(client *mongo.Client, hero Hero) interface{} {
	collection := client.Database("civilact").Collection("heros")
	insertResult, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	return insertResult.InsertedID
}

func RemoveOneHero(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heros")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Hero", err)
	}
	return deleteResult.DeletedCount
}

func UpdateHero(client *mongo.Client, updatedData interface{}, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heros")
	atualizacao := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
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

	/*
		heroes := ReturnAllHeroes(c, bson.M{})
		//heroes := ReturnAllHeroes(c, bson.M{"signed": true})
		for _, hero := range heroes {
			log.Println(hero.Name, hero.Alias, hero.Signed)
		}

		hero := ReturnOneHero(c, bson.M{"name": "Vision"})
		log.Println(hero.Name, hero.Alias, hero.Signed)


		var hero = Hero{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: true}
		insertedID := InsertNewHero(c, hero)
		log.Println(insertedID)
		hero = ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
		log.Println(hero.Name, hero.Alias, hero.Signed)


		heroesRemoved := RemoveOneHero(c, bson.M{"Alias": "Doctor Strange"})
		log.Println("Heroes removed count:", heroesRemoved)
		var hero = ReturnOneHero(c, bson.M{"Alias": "Doctor Strange"})
		log.Println("Is Hero empty?", hero == Hero{})


		heroesUpdated := UpdateHero(c, bson.M{"signed": false}, bson.M{"name": "Steve Rodgers"})
		log.Println("Heroes updated count:", heroesUpdated)
		var hero = ReturnOneHero(c, bson.M{"alias": "Captian America"})
		log.Println(hero.Name, hero.Alias, hero.Signed)
	*/
}

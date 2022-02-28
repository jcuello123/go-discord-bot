package db

import (
	"go-discord-bot/config"
	"go-discord-bot/constants"
	"log"

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
	coll *mongo.Collection
) 


func Connect(){
	config.GetDbURI()	

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB_URI))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged DB")

	coll = client.Database("go-discord-bot").Collection("maps")
}

func GetAllMapsAsBson() []bson.D{
	var bsonDocs []bson.D
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		
		bsonDocs = append(bsonDocs, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return bsonDocs
}

func CompleteMap(mapName string) string {
	if !constants.MapExists(mapName){
		return fmt.Sprintf("%s doesn't exist.", mapName)
	}

	filter := bson.D{{"name", mapName}}
	update := bson.D{{"$set", bson.D{{"completed", true}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("Error when marking %s as complete.", mapName)
	}

	return fmt.Sprintf("%s has been marked as complete.", mapName)
}
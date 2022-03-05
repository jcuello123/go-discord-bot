package db

import (
	"errors"
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


func Connect() error{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetDbURI()))

	if err != nil {
		return err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("Successfully connected and pinged DB")

	coll = client.Database("go-discord-bot").Collection("maps")
	return nil
}

func GetAllMapsAsBson() ([]bson.D, error){
	var bsonDocs []bson.D
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		
		bsonDocs = append(bsonDocs, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return bsonDocs, nil
}

func UpdateMapComplete(mapName string, completed bool) error {
	if !constants.MapExists(mapName){
		return errors.New(mapName + " doesn't exist") 
	}

	filter := bson.D{{"name", mapName}}
	update := bson.D{{"$set", bson.D{{"completed", completed}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		errMsg := fmt.Sprintf("Error when updating %s.", mapName)
		return errors.New(errMsg) 
	}

	return nil 
}
package db

import (
	"go-discord-bot/config"
	// "go-discord-bot/maps"

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
)


func Connect(){
	config.GetDbURI()	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB_URI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Print("Successfully connected and pinged DB")

	// coll := client.Database("go-discord-bot").Collection("maps")
	// filter := bson.D{}
	
	// var result bson.D
	// err = coll.FindOne(context.TODO(), filter).Decode(&result)
	// var mappy maps.ZombieMap

	// doc, _:= bson.Marshal(bson.D{result[1]})
	// err = bson.Unmarshal(doc, &mappy)

	// fmt.Println("GOT MAP: ", mappy.Name)
}
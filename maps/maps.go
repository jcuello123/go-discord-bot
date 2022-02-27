package maps

import (
	"go-discord-bot/constants"
	"go-discord-bot/db"

	"log"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ZombieMap struct {
	Name string
	Completed bool
}

func GetRandMap() string {
	rand.Seed(time.Now().UnixNano())
	zombieMap := constants.ZombieMapsArr[rand.Intn(len(constants.ZombieMapsArr))]
	return strings.ToUpper(zombieMap) 
}

func bsonToZombieMaps() []ZombieMap{
	var zombieMaps []ZombieMap
	bsonDocs := db.GetAllMapsAsBson()

	for _, result := range bsonDocs {
		nameAsDoc, err := bson.Marshal(bson.D{result[1]})
		if err != nil {
			log.Fatal(err)
		}

		// TODO: figure out how to marshal all the properties as once
		completedAsDoc, err := bson.Marshal(bson.D{result[2]})
		if err != nil {
			log.Fatal(err)
		}

		var zombieMap ZombieMap

		err = bson.Unmarshal(nameAsDoc, &zombieMap)
		if err != nil {
			log.Fatal(err)
		}

		err = bson.Unmarshal(completedAsDoc, &zombieMap)
		if err != nil {
			log.Fatal(err)
		}

		zombieMaps = append(zombieMaps, zombieMap)
	}

	return zombieMaps
}

func FormattedMaps() string {
	var result strings.Builder
	zombieMaps := bsonToZombieMaps()

	for _, zombieMap := range zombieMaps {
		result.WriteString(zombieMap.Name)
		result.WriteString(" ")
		
		if zombieMap.Completed {
			result.WriteString(constants.CHECK_MARK)
		} else {
			result.WriteString(constants.X)
		}
		
		result.WriteString("\n")
	}

	return result.String()
}

func CompleteMap(args []string) string {
	var mapName strings.Builder
	lastElement := len(args) - 1
	for i, arg := range args[1:] {
		mapName.WriteString(arg)
		if i < lastElement {
			mapName.WriteString(" ")
		}
	} 

	return db.CompleteMap(mapName.String())
}
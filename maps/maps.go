package maps

import (
	"fmt"
	"go-discord-bot/constants"
	"go-discord-bot/db"
	"go-discord-bot/emoji"

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

	emoji, err := emoji.ZMapToEmoji(zombieMap)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	msg := fmt.Sprintf("%s %s %s", emoji, strings.ToUpper(zombieMap), emoji)
	return msg 
}

func bsonToZombieMaps() ([]ZombieMap, error){
	var zombieMaps []ZombieMap

	bsonDocs, err := db.GetAllMapsAsBson()
	if err != nil {
		return nil, err
	}

	for _, result := range bsonDocs {
		nameAsDoc, err := bson.Marshal(bson.D{result[1]})
		if err != nil {
			return nil, err
		}

		// TODO: figure out how to marshal all the properties as once
		completedAsDoc, err := bson.Marshal(bson.D{result[2]})
		if err != nil {
			return nil, err
		}

		var zombieMap ZombieMap

		err = bson.Unmarshal(nameAsDoc, &zombieMap)
		if err != nil {
			return nil, err
		}

		err = bson.Unmarshal(completedAsDoc, &zombieMap)
		if err != nil {
			return nil, err
		}

		zombieMaps = append(zombieMaps, zombieMap)
	}

	return zombieMaps, nil
}

func FormattedMaps() string {
	var result strings.Builder
	
	zombieMaps, err := bsonToZombieMaps()
	if err != nil {
		return ""
	}

	for _, zombieMap := range zombieMaps {
		result.WriteString(zombieMap.Name)
		result.WriteString(" ")
		
		if zombieMap.Completed {
			result.WriteString(emoji.CHECK_MARK)
		} else {
			result.WriteString(emoji.X)
		}
		
		result.WriteString("\n")
	}

	return result.String()
}

func UpdateMapComplete(args []string, completed bool) error{
	var mapName strings.Builder
	lastElement := len(args) - 1
	for i, arg := range args {
		mapName.WriteString(arg)
		if i < lastElement {
			mapName.WriteString(" ")
		}
	} 

	err := db.UpdateMapComplete(mapName.String(), completed)
	if err != nil {
		return err
	}
	return nil
}
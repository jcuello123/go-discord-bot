package maps

import (
	"go-discord-bot/constants"
	"go-discord-bot/db"
	"log"
	"testing"
)

func init() {
	err := db.Connect()
	if err != nil {
		log.Println("Error when connecting to local db for testing.", err)
	}
}

func TestGetRandomMap(t *testing.T) {
	zMap := GetRandMap()
	if zMap == "" {
		t.Errorf("Expected a zombie map but received empty")
	}
}

func TestBsonToZombieMaps(t *testing.T) {
	zMaps, err := bsonToZombieMaps()
	if err != nil {
		t.Errorf("Expected error to be nil but received %s", err)
	}
	if len(zMaps) != len(constants.ZombieMapsArr) {
		t.Errorf("Expected length %d but received %d", len(constants.ZombieMapsArr), len(zMaps))
	}
}

func TestFormattedMaps(t *testing.T) {
	formattedMaps := FormattedMaps()
	if formattedMaps == "" {
		t.Errorf("Expected maps to be formatted but received empty string")
	}
}

func TestUpdateMapComplete(t *testing.T) {
	args := []string{"nacht", "der", "untoten"}
	
	err := UpdateMapComplete(args, false)
	if err != nil {
		t.Errorf("Expected nil error but received %s", err)
	}

	args = []string{"invalid"}

	err = UpdateMapComplete(args, false)
	if err == nil {
		t.Errorf("Expected error but received nil")
	}
}
package db

import (
	"go-discord-bot/constants"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Errorf("Expecteed nil but received %s", err)
	}
}

func TestGetAllMapsAsBson(t *testing.T) {
	mapsAsBson, err := GetAllMapsAsBson()
	if err != nil {
		t.Errorf("Expected nil but received %s", err)
	}
	if len(mapsAsBson) != len(constants.ZombieMapsArr){
		t.Errorf("Expected length %d but received %d", len(constants.ZombieMapsArr), len(mapsAsBson))
	}
}

func TestUpdateMapComplete(t *testing.T) {
	err := UpdateMapComplete("verruckt", true)
	if err != nil {
		t.Errorf("Expected nil but received %s", err)
	}

	err = UpdateMapComplete("invalidMap", false)
	if err == nil {
		t.Errorf("Expected map to not exist error but received nil")
	}
}
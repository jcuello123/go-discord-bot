package constants

import (
	"go-discord-bot/types"

	"fmt"
	"log"
)

var (
    ZombieMapsArr = []string{"nacht der untoten", "verruckt", "shang ri la", "moon", "origins", "shi no numa", "shadows of evil", "der riese", "ascension", "kino der toten"}
	ZombieMapsSet = createMapsSet() 
	empty types.Void 
) 

func createMapsSet() map[string]types.Void{
	temp := make(map[string]types.Void)
	for _, zMap := range ZombieMapsArr{
		temp[zMap] = empty
	}
	return temp
}

func MapExists(mapName string) bool {
	_, exists := ZombieMapsSet[mapName]
	if !exists {
		errMsg := fmt.Sprintf("%s doesn't exist", mapName)
		log.Println(errMsg)
	}
	return exists
}
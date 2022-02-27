package constants

import (
	"go-discord-bot/types"
)

var (
    ZombieMapsArr = []string{"nacht der untoten", "verruckt", "shang ri la", "moon", "origins", "shi no numa", "shadow of evil", "der riese", "ascension", "kino der toten"}
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
	return exists
}
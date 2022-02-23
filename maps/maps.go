package maps

import (
	"math/rand"
	"time"
	"strings"
)

var zMaps = []string{"nacht der untoten", "verruckt", "shang ri la", "moon", "origins", "shi no numa", "shadow of evil", "der riese", "ascension", "kino der toten"}

type ZombieMap struct {
	Name string
	Completed bool
}


func GetRandMap() string {
	rand.Seed(time.Now().UnixNano())
	zombieMap := zMaps[rand.Intn(len(zMaps))]
	return strings.ToUpper(zombieMap) 
}
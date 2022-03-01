package main

import (
	"fmt"
	"go-discord-bot/bot"
	"go-discord-bot/config"
)

func main(){
	if err := config.ReadConfig(config.File); err != nil{
		fmt.Println(err.Error())	
		return
	}

	bot.Start()
	
	<-make(chan struct{})
	return
}
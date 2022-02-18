package bot

import (
	"fmt"
	"go-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session

func Start(){
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	if firstChar := m.Content[0:1]; firstChar != config.BotPrefix || m.Author.ID == botId {
		return
	}
	
	command := m.Content[1:]

	if command == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
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
	fmt.Println("IN MESSAGE HANDLER")

	if m.Author.ID == botId {
		return
	}

	fmt.Println("m.Content:",m.Content)

	if m.Content == "ping" {
	fmt.Println("PING")
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
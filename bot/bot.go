package bot

import (
	"fmt"
	"log"
	"strings"

	"go-discord-bot/commands"
	"go-discord-bot/config"
	"go-discord-bot/db"

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

	err = db.Connect()
	if err != nil {
		panic(err)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	if !strings.HasPrefix(m.Content, config.BotPrefix) || m.Author.ID == botId {
		return
	}
	
	args := strings.Split(m.Content, " ")
	commandWithPrefix := args[0]
	if len(commandWithPrefix) == 1 {
		return
	} 

	command := strings.Replace(commandWithPrefix, config.BotPrefix, "", 1)
	
	err := commands.Execute(command, args[1:], s, m.ChannelID)
	if err != nil {
		log.Println(err)
	}
}
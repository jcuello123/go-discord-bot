package bot

import (
	"fmt"
	"go-discord-bot/config"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session
var maps = []string{"nacht der untoten", "verruckt", "shang ri la", "moon", "origins", "shi no numa", "shadow of evil", "der riese"}

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
	if !strings.HasPrefix(m.Content, config.BotPrefix) || m.Author.ID == botId {
		return
	}
	
	args := strings.Split(m.Content, " ")
	commandWithPrefix := args[0]
	if len(commandWithPrefix) == 1 {
		return
	} 

	command := strings.Replace(commandWithPrefix, config.BotPrefix, "", 1)

	if command == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if command == "maps" {
		_, err := s.ChannelMessageSend(m.ChannelID, getRandMap())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if len(args) > 2 && command == "random" {
		_, err := s.ChannelMessageSend(m.ChannelID, getRandItem(args[1:]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func getRandMap() string {
	rand.Seed(time.Now().UnixNano())
	zombieMap := maps[rand.Intn(len(maps))]
	return strings.ToUpper(zombieMap) 
}

func getRandItem(args[] string) string {
	rand.Seed(time.Now().UnixNano())
	return args[rand.Intn(len(args))]
}
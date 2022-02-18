package bot

import (
	"fmt"
	"go-discord-bot/config"
	"math/rand"
    "time"
	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session
var maps = []string{"nacht der untoten", "verruckt", "shang ri la", "moon", "origins", "shi no numa", "shadow of evil", "der riese"}

func init(){
	rand.Seed(time.Now().UnixNano())
}

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

	if command == "maps" {
		_, err := s.ChannelMessageSend(m.ChannelID, getRandMap())
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getRandMap() string{
	return maps[rand.Intn(len(maps))]
}
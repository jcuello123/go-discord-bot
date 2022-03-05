package commands

import (
	"math/rand"
	"fmt"
	"time"
	"errors"

	"go-discord-bot/maps"
	"go-discord-bot/types"

	"github.com/bwmarrin/discordgo"
)

type ping struct {}
type randMap struct {}
type random struct {}
type completed struct {}
type completeMap struct {}
type unCompleteMap struct {}

type command interface{
	execute(args []string)
}

var (
	allCommands = make(map[string]types.Void)
	empty types.Void
	session *discordgo.Session
	msgCreate *discordgo.MessageCreate
) 

func init() {
	allCommands["ping"] = empty
	allCommands["randmap"] = empty
	allCommands["random"] = empty
	allCommands["completed"] = empty
	allCommands["completemap"] = empty
	allCommands["uncompletemap"] = empty
}

func Execute(command string, args []string,  s *discordgo.Session, m *discordgo.MessageCreate) {
	_, exists := allCommands[command]
	if !exists {
		return
	}

	session = s
	msgCreate = m

	cmd, err := createCommand(command)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd.execute(args)
}

func createCommand(cmd string) (command, error) {
	if cmd == "ping" {
		var p ping 
		return p, nil
	}

	if cmd == "randmap" {
		var rm randMap
		return rm, nil
	}

	if cmd == "random" {
		var r random
		return r, nil
	}

	if cmd == "completed" {
		var c completed 
		return c, nil
	}

	if cmd == "completemap" {
		var cm completeMap 
		return cm, nil	
	}

	if cmd == "uncompletemap" {
		var ucm unCompleteMap 
		return ucm, nil	
	}

	errMsg := fmt.Sprintf("The '%s' command couldn't be created.", cmd)
	return nil, errors.New(errMsg)
}

func sendMessage(message string) {
	if message == "" {
		return
	}

	_, err := session.ChannelMessageSend(msgCreate.ChannelID, message)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (p ping) execute(args []string) {
	sendMessage("pong")
}

func (rm randMap) execute(args []string) {
	sendMessage(maps.GetRandMap())
}

func (r random) execute(args []string) {
	sendMessage(getRandItem(args))
}

func (c completed) execute(args []string) {
	sendMessage(maps.FormattedMaps())
}

func (cm completeMap) execute(args []string) {
	err := maps.UpdateMapComplete(args, true)
	if err != nil {
		sendMessage(err.Error())
	} else {
		sendMessage(maps.FormattedMaps())
	}
}

func (ucm unCompleteMap) execute(args []string) {
	err := maps.UpdateMapComplete(args, false)
	if err != nil {
		sendMessage(err.Error())
	} else {
		sendMessage(maps.FormattedMaps())
	}
}

func getRandItem(args[] string) string {
	rand.Seed(time.Now().UnixNano())
	return args[rand.Intn(len(args))]
}
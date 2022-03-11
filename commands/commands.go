package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"go-discord-bot/battleship"
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
type battleShip struct {}

type command interface{
	execute(args []string) error
}

type discordSession interface {
	ChannelMessageSend(channelID string, message string) (*discordgo.Message, error)
}

var (
	allCommands = make(map[string]types.Void)
	empty types.Void
	session discordSession 
	channelID string
) 

func init() {
	allCommands["ping"] = empty
	allCommands["randmap"] = empty
	allCommands["random"] = empty
	allCommands["completed"] = empty
	allCommands["completemap"] = empty
	allCommands["uncompletemap"] = empty
	allCommands["bs"] = empty
}

// TODO: fix returning error from every method. just log the error in that method

func Execute(command string, args []string,  s discordSession, discordChannelID string) error {
	if !commandExists(command) {
		return errors.New("Invalid command: " + command)
	}

	session = s
	channelID = discordChannelID 

	cmd, err := createCommand(command)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cmd.execute(args)
	return nil
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

	if cmd == "bs" {
		var bs battleShip
		return bs, nil 
	}

	errMsg := fmt.Sprintf("The '%s' command couldn't be created.", cmd)
	return nil, errors.New(errMsg)
}

func sendMessage(message string) error {
	if message == "" {
		return errors.New("Empty message")
	}

	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		return err
	}

	return nil
}

func (p ping) execute(args []string) error {
	return sendMessage("pong")
}

func (rm randMap) execute(args []string) error {
	return sendMessage(maps.GetRandMap())
}

func (r random) execute(args []string) error {
	return sendMessage(getRandItem(args))
}

func (c completed) execute(args []string) error {
	return sendMessage(maps.FormattedMaps())
}

func (cm completeMap) execute(args []string) error {
	err := maps.UpdateMapComplete(args, true)
	if err != nil {
		return sendMessage(err.Error())
	} else {
		return sendMessage(maps.FormattedMaps())
	}
}

func (ucm unCompleteMap) execute(args []string) error {
	err := maps.UpdateMapComplete(args, false)
	if err != nil {
		return sendMessage(err.Error())
	} else {
		return sendMessage(maps.FormattedMaps())
	}
}

func (bs battleShip) execute(args []string) error {
	if len(args) == 0 {
		return errors.New("Battleship command takes at least one more argument.")
	}

	subcmd := args[0]
	
	if subcmd == "help" {
		return sendMessage(battleship.Help)
	} else if subcmd == "start" {
		battleship.Start()
		return sendMessage(battleship.GetBoardAsString())
	} else if subcmd == "shoot" {
		msg, err := battleship.Shoot(args)
		if err != nil {
			if err.Error() == "Max attempts reached. Game over." {
				sendMessage(battleship.GetBoardAsString())
			}
			return sendMessage(err.Error())
		}
		sendMessage(battleship.GetBoardAsString())
		sendMessage(msg)
	}

	return errors.New("Invalid use of bs command") 
}

func getRandItem(args[] string) string {
	rand.Seed(time.Now().UnixNano())
	return args[rand.Intn(len(args))]
}

func commandExists(command string) bool {
	_, exists := allCommands[command]
	return exists
}
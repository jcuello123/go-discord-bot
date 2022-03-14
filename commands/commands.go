package commands

import (
	"errors"
	"log"
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
		return errors.New("Invalid command: %s" + command)
	}

	session = s
	channelID = discordChannelID 

	cmd := createCommand(command)
	
	err := cmd.execute(args)
	return err	
}

func createCommand(cmd string) command {
	if cmd == "ping" {
		var p ping 
		return p
	}

	if cmd == "randmap" {
		var rm randMap
		return rm
	}

	if cmd == "random" {
		var r random
		return r
	}

	if cmd == "completed" {
		var c completed 
		return c
	}

	if cmd == "completemap" {
		var cm completeMap 
		return cm	
	}

	if cmd == "uncompletemap" {
		var ucm unCompleteMap 
		return ucm	
	}

	if cmd == "bs" {
		var bs battleShip
		return bs 
	}

	log.Printf("The '%s' command couldn't be created.", cmd)
	return nil
}

func sendMessage(message string) error {
	if message == "" {
		return errors.New("Message is empty.")
	}

	_, err := session.ChannelMessageSend(channelID, message)
	return err
}

func (p ping) execute(args []string) error{
	err := sendMessage("pong")
	return err
}

func (rm randMap) execute(args []string) error {
	err := sendMessage(maps.GetRandMap())
	return err
}

func (r random) execute(args []string) error {
	err := sendMessage(getRandItem(args))
	return err
}

func (c completed) execute(args []string) error {
	err := sendMessage(maps.FormattedMaps())
	return err
}

func (cm completeMap) execute(args []string) error {
	updateErr := maps.UpdateMapComplete(args, true)
	var err error
	if updateErr != nil {
		err = sendMessage(updateErr.Error())
	} else {
		err =sendMessage(maps.FormattedMaps())
	}
	return err
}

func (ucm unCompleteMap) execute(args []string) error {
	updateErr:= maps.UpdateMapComplete(args, false)
	var err error
	if updateErr != nil {
		err = sendMessage(updateErr.Error())
	} else {
		err = sendMessage(maps.FormattedMaps())
	}
	return err
}

func (bs battleShip) execute(args []string) error {
	if len(args) == 0 {
		return errors.New("Battleship command takes at least one more argument.")
	}

	subcmd := args[0]
	var err error
	
	if subcmd == "help" {
		err = sendMessage(battleship.Help)
	} else if subcmd == "start" {
		battleship.Start()
		err = sendMessage(battleship.GetBoardAsString())
	} else if subcmd == "shoot" {
		msg, shootErr:= battleship.Shoot(args[1:])
		if shootErr != nil {
			if shootErr.Error() == "Max attempts reached. Game over." {
				err = sendMessage(battleship.GetBoardAsString())
				if err != nil {
					return err
				}
			}
			err = sendMessage(shootErr.Error())
			if err != nil {
				return err
			}
		}
		err = sendMessage(battleship.GetBoardAsString())
		if err != nil {
			return err
		}
		if msg != "" {
			err = sendMessage(msg)
			if err != nil {
				return err
			}
		}
	} else {
		err = sendMessage("Invalid use of !bs shoot.")
		if err != nil {
			return err
		}
	}

	return err
}

func getRandItem(args[] string) string {
	rand.Seed(time.Now().UnixNano())
	return args[rand.Intn(len(args))]
}

func commandExists(command string) bool {
	_, exists := allCommands[command]
	return exists
}
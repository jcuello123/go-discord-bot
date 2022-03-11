package commands

import (
	"go-discord-bot/db"
	"log"
	"testing"

	"github.com/bwmarrin/discordgo"
)

var sessionMock discordSessionMock

type discordSessionMock struct {
}

func init() {
	err := db.Connect()
	if err != nil {
		log.Fatal("Error when connecting to local db for testing.", err)
	}
}

func (d discordSessionMock) ChannelMessageSend(channelID string, message string) (*discordgo.Message, error) {
	return nil, nil
}

func TestExecuteInvalidCommand(t *testing.T) {
	args := []string{}

	err := Execute("invalid", args, sessionMock, "1")
	if err == nil {
		t.Fatalf("Expected invalid command error but received: %s", err)
	}
}

func TestExecutePingCommand(t *testing.T) {
	args := []string{}

	err := Execute("ping", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteRandMapCommand(t *testing.T) {
	args := []string{}

	err := Execute("randmap", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteRandomCommand(t *testing.T) {
	args := []string{"1", "2", "3"}

	err := Execute("random", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteCompletedCommand(t *testing.T) {
	args := []string{}

	err := Execute("completed", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteCompleteMapCommand(t *testing.T) {
	args := []string{"der riese"}

	err := Execute("completemap", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteUnCompleteMapCommand(t *testing.T) {
	args := []string{"verruckt"}

	err := Execute("uncompletemap", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteBSHelpCommand(t *testing.T) {
	args := []string{"help"}
	
	err := Execute("bs", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteBSStartCommand(t *testing.T) {
	args := []string{"start"}
	
	err := Execute("bs", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}

func TestExecuteBSShootCommand(t *testing.T) {
	args := []string{"shoot", "0", "0"}
	
	err := Execute("bs", args, sessionMock, "1")
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err)
	}
}
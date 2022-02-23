package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
)

var (
	Token string
	BotPrefix string
	config *configStruct
	DB_URI string
	localDBURI string = "mongodb://localhost:27017"
)

type configStruct struct{
	Token string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}


func ReadConfig() error{
	fmt.Println("Reading config file..")

	if err := readConfigFromEnv(); err == nil {
		return nil
	}

	file, err := ioutil.ReadFile("./config.json")

	if err != nil{
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)

	if err != nil{
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}

func readConfigFromEnv() error {
	token := os.Getenv("TOKEN")
	if token == "" {
		return errors.New("Bot token is missing.")
	}

	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		return errors.New("Bot prefix is missing.")
	}

	Token = token
	BotPrefix = botPrefix
	fmt.Println("Read config from env successfully.")

	return nil
}

func GetDbURI() {
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		DB_URI = localDBURI
		fmt.Println("Defaulting to local db URI.")
		return
	}

	DB_URI = dbURI
	fmt.Println("Read DB_URI from env successfully.")
}
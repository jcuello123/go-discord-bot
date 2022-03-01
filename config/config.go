package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	Token string
	BotPrefix string
	config *configStruct
	localDBURI string = "mongodb://localhost:27017"
	File string = "./config/config.json"
)

type configStruct struct{
	Token string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig(configFile string) error{
	fmt.Println("Reading config file..")

	if err := readConfigFromEnv(); err == nil {
		return nil
	}

	file, err := ioutil.ReadFile(configFile)

	if err != nil{
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(config)

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

func GetDbURI() string{
	dbURIFromEnv:= os.Getenv("DB_URI")
	if dbURIFromEnv == "" {
		fmt.Println("Defaulting to local db URI.")
		return localDBURI
	}

	fmt.Println("Read DB_URI from env successfully.")
	return dbURIFromEnv
}
package config

import (
	"os"
	"testing"
)

func TestGetDbURIFromEnv(t *testing.T){
	testDbURI := "DB_URI"
	testDbURIVal := "localhost"

	os.Setenv(testDbURI, testDbURIVal)

	dbURIFromEnv := GetDbURI()
	if dbURIFromEnv != testDbURIVal {
		os.Unsetenv(testDbURI)
		t.Errorf("Expected %s but received %s from env", testDbURIVal, dbURIFromEnv)
	}

	os.Unsetenv(testDbURI)
}

func TestGetDbURILocal(t *testing.T) {
	dbURI := GetDbURI()
	if dbURI != localDBURI {
		t.Errorf("Expected %s but received %s from local", localDBURI, dbURI)
	}
}

func TestReadConfigProperEnvVariablesSet(t *testing.T) {
	os.Setenv("TOKEN", "token")
	os.Setenv("BOT_PREFIX", "prefix")

	err := ReadConfig("config_test.json")
	if err != nil {
		t.Errorf("Expected nil but received %s", err)
	}
	
	os.Unsetenv("TOKEN")
	os.Unsetenv("BOT_PREFIX")
}

func TestReadConfigLocal(t *testing.T) {
	err := ReadConfig("config_test.json")
	if err != nil {
		t.Errorf("Expected nil but received %s", err)
	}
	if config.Token != "token" {
		t.Errorf("Expected token but received %s", config.Token)
	}
	if config.BotPrefix != "prefix" {
		t.Errorf("Expected prefix but received %s", config.BotPrefix)
	}
}

func TestReadConfigLocalWithOnlyTokenSet(t *testing.T) {
	os.Setenv("TOKEN", "token")
	err := ReadConfig("config_test.json")

	if err != nil {
		t.Errorf("Expected nil but received %s", err)
	}
	if  config.Token != "token" {
		t.Errorf("Expected token but received %s", config.Token)
	}
	if config.BotPrefix != "prefix" {
		t.Errorf("Expected prefix but received %s", config.BotPrefix)
	}
}

func TestReadConfigInvalidFile(t *testing.T) {
	err := ReadConfig("invalid_file.json")
	if err == nil {
		t.Errorf("Expected error but received nil")
	}
}
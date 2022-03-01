package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDbURIFromEnv(t *testing.T){
	testDbURI := "DB_URI"
	testDbURIVal := "localhost"

	os.Setenv(testDbURI, testDbURIVal)

	dbURIFromEnv := GetDbURI()

	assert.Equal(t, testDbURIVal, dbURIFromEnv, "Should retrieve DB URI from env")

	os.Unsetenv(testDbURI)
}

func TestGetDbURILocal(t *testing.T) {
	dbURI := GetDbURI()
	assert.Equal(t, localDBURI, dbURI, "Should default to local DB URI when env variable isnt present")
}

func TestReadConfigProperEnvVariablesSet(t *testing.T) {
	os.Setenv("TOKEN", "token")
	os.Setenv("BOT_PREFIX", "prefix")

	err := ReadConfig("config_test.json")

	assert.Nil(t, err)
	
	os.Unsetenv("TOKEN")
	os.Unsetenv("BOT_PREFIX")
}

func TestReadConfigLocal(t *testing.T) {
	err := ReadConfig("config_test.json")

	assert.Nil(t, err)
	assert.Equal(t, "token", config.Token)
	assert.Equal(t, "prefix", config.BotPrefix)
}

func TestReadConfigLocalWithOnlyTokenSet(t *testing.T) {
	os.Setenv("TOKEN", "token")
	err := ReadConfig("config_test.json")

	assert.Nil(t, err)
	assert.Equal(t, "token", config.Token)
	assert.Equal(t, "prefix", config.BotPrefix)
}

func TestReadConfigInvalidFile(t *testing.T) {
	err := ReadConfig("invalid_file.json")
	
	assert.Error(t, err)
}
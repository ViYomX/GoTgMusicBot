package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIID         int
	APIHash       string
	BotToken      string
	StringSession string
	StartImageUrl string
)

func init() {
	godotenv.Load()

	var err error
	APIID, err = strconv.Atoi(getenv("API_ID", "0"))
	if err != nil {
		panic("Invalid API_ID: " + getenv("API_ID", ""))
	}

	APIHash = getenv("API_HASH", "")
	BotToken = getenv("BOT_TOKEN", "")
	StringSession = getenv("STRING_SESSION", "")
	StartImageUrl = getenv("START_IMG_URL", "")
}

func getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
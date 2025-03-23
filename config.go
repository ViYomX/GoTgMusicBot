package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APIID         string
	APIHash       string
	BotToken      string
	StringSession string
)

func init() {
	godotenv.Load()

	APIID = Getenv("API_ID", "")
	APIHash = Getenv("API_HASH", "")
	BotToken = Getenv("BOT_TOKEN", "")
	StringSession = Getenv("STRING_SESSION", "")
}

func Getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

package config

import (
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

	APIID = Atoi(getenv("API_ID", ""))
	APIHash = getenv("API_HASH", "")
	BotToken = getenv("BOT_TOKEN", "")
	StringSession = getenv("STRING_SESSION", "")
}

func getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

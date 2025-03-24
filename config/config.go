package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIID         int32
	APIHash       string
	BotToken      string
	StartImageUrl string
)

func init() {
	godotenv.Load()

	apiID, err := strconv.Atoi(getenv("API_ID", "0"))
	if err != nil {
		panic("Invalid API_ID: " + getenv("API_ID", "12380656"))
	}
	APIID = int32(apiID)

	APIHash = getenv("API_HASH", "d927c13beaaf5110f25c505b7c071273")
	BotToken = getenv("BOT_TOKEN", "")
	StartImageUrl = getenv("START_IMG_URL", "https://graph.org/file/d05729b62d49be693f360-7a214538921cf8c335.jpg")
}

func getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

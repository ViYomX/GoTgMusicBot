package utils

import (
	"log"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

var (
	BOT_TOKEN      = ""
	API_KEY        int
	API_HASH       string
	STRING_SESSION string
)

func InitEnv() {
	dotenv.Load()
	BOT_TOKEN = getEnv("BOT_TOKEN")
	API_KEY, _ = strconv.Atoi(getEnv("API_KEY"))
	API_HASH = getEnv("API_HASH")
	STRING_SESSION = getEnv("STRING_SESSION")
}

func getEnv(key string) string {
	key, ok := os.LookupEnv(key)
	if key == "" || !ok {
		log.Fatalf("Environment variable %s not set", key)
	}

	return key
}

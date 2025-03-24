package main

import (
	"strings"

	"main/config"
	"main/modules"

	tg "github.com/amarnathcjd/gogram/telegram"
)
func main() {

	bot, _ := tg.NewClient(tg.ClientConfig{
		AppID:    config.APIID,
		AppHash:  config.APIHash,
		Session:  "bot.dat",
		LogLevel: tg.LogInfo,
		Cache: tg.NewCache("bot.cache", &tg.CacheConfig{
			LogLevel: tg.LogInfo,
		}),
	})

	bot.LoginBot(config.BotToken)
	bot.On("message:/start", modules.StartHandler)

	bot.Idle()
}

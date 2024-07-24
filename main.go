package main

//#cgo LDFLAGS: -L . -lntgcalls -Wl,-rpath=./

import "C"
import (
	"fmt"
	"main/utils"

	"github.com/amarnathcjd/gogram/telegram"
)

func main() {
	utils.InitEnv()
	ub, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:         int32(utils.API_KEY),
		AppHash:       utils.API_HASH,
		StringSession: utils.STRING_SESSION,
		MemorySession: true,
	})

	bot, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:   int32(utils.API_KEY),
		AppHash: utils.API_HASH,
	})

	bot.Conn()
	bot.LoginBot(utils.BOT_TOKEN)

	setupCallsCore()
	ub.Log.Info("ntgcalls - core - started")

	url := "https://envs.sh/trq.m4a" // audio file url

	media := MediaDescription{
		Audio: &AudioDescription{
			InputMode:     InputModeShell,
			SampleRate:    128000,
			BitsPerSample: 16,
			ChannelCount:  2,
			Input:         fmt.Sprintf("ffmpeg -i %s -loglevel panic -f s16le -ac 2 -ar 128k pipe:1", url), // ffmpeg command to convert audio to s16le format and pipe it to stdout
		},
	}

	// media.Video = &VideoDescription{
	// 	InputMode: InputModeShell,
	// 	Input:     fmt.Sprintf("ffmpeg -i %s -loglevel panic -f rawvideo -r 24 -pix_fmt yuv420p -vf scale=1280:720 pipe:1", video),
	// 	Width:     1280,
	// 	Height:    720,
	// 	Fps:       24,
	// }

	joinGroupCall(ub, call, "@rosexchat", media)

}

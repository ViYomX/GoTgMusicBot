package main

//#cgo LDFLAGS: -L . -lntgcalls -Wl,-rpath=./
import "C"

import (
	"fmt"
	"os"
	"strings"

	"main/ntgcalls"

	tg "github.com/amarnathcjd/gogram/telegram"
	dotenv "github.com/joho/godotenv"
)

var (
	caller *ntgcalls.Client
	client *tg.Client
)

func main() {
	dotenv.Load()
	caller = ntgcalls.NTgCalls()
	defer caller.Free()
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:    Atoi(os.Getenv("API_ID")),
		AppHash:  Atoi(os.Getenv("API_HASH")),
		Session:  os.Getenv("STRING_SESSION"),
		LogLevel: tg.LogInfo,
	})
	client.Start()

	bot, _ := tg.NewClient(tg.ClientConfig{
		AppID:    Atoi(os.Getenv("API_ID")),
		AppHash:  Atoi(os.Getenv("API_HASH")),
		Session:  "bot.dat",
		LogLevel: tg.LogInfo,
		Cache: tg.NewCache("bot.cache", &tg.CacheConfig{
			LogLevel: tg.LogInfo,
		}),
	})

	bot.LoginBot(os.Getenv("BOT_TOKEN"))
	bot.On("message:/start", StartHandler)
	bot.On("message:!play", playHandler)

	client.Idle()
}

func StartHandler(m *tg.NewMessage) error {
	fmt.Println("Bot started")
	m.Reply("VCPlayBot is Active!")
	return nil
}

func playHandler(m *tg.NewMessage) error {
	if !m.IsReply() {
		m.Reply("Reply to an audio file to play it!")
		return nil
	}

	r, err := m.GetReplyMessage()
	if r.Audio() == nil || err != nil {
		m.Reply("Reply to an audio file to play it!")
		return nil
	}

	msg, _ := m.Respond("<code>Downloading...</code>")

	file, _ := r.Download()
	msg.Edit("<code>Converting...</code>")
	convertedFile, err := convertToSle3(file)
	if err != nil {
		m.Reply("Error converting file")
		return nil
	}

	msg.Edit("<code>Playing...</code>")

	call, err := caller.CreateCall(m.ChatID(), ntgcalls.MediaDescription{
		Microphone: &ntgcalls.AudioDescription{
			MediaSource:  ntgcalls.MediaSourceFile,
			SampleRate:   128000,
			ChannelCount: 2,
			Input:        convertedFile,
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "cannot be initialized more") {
			caller.Stop(m.ChatID())
			call, _ = caller.CreateCall(m.ChatID(), ntgcalls.MediaDescription{
				Microphone: &ntgcalls.AudioDescription{
					MediaSource:  ntgcalls.MediaSourceFile,
					SampleRate:   128000,
					ChannelCount: 2,
					Input:        convertedFile,
				},
			})
		} else {
			m.Reply("Error playing file: " + err.Error())
			return nil
		}
	}

	if m.Channel != nil {
		if m.Channel.Username != "" {
			channel, _ := client.GetSendableChannel(m.Channel.Username)

			fullChatRaw, _ := client.ChannelsGetFullChannel(
				&tg.InputChannelObj{
					ChannelID:  channel.(*tg.InputChannelObj).ChannelID,
					AccessHash: channel.(*tg.InputChannelObj).AccessHash,
				},
			)
			fullChat := fullChatRaw.FullChat.(*tg.ChannelFull)
			callResRaw, _ := client.PhoneJoinGroupCall(
				&tg.PhoneJoinGroupCallParams{
					Muted:        false,
					VideoStopped: true,
					Call:         fullChat.Call,
					Params: &tg.DataJson{
						Data: call,
					},
					JoinAs: &tg.InputPeerUser{
						UserID:     client.Me().ID,
						AccessHash: client.Me().AccessHash,
					},
				},
			)
			callRes := callResRaw.(*tg.UpdatesObj)
			for _, update := range callRes.Updates {
				switch upd := update.(type) {
				case *tg.UpdateGroupCallConnection:
					_ = caller.Connect(m.ChatID(), upd.Params.Data, false)
				}
			}
		}
	}

	return nil
}

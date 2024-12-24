package main

//#cgo LDFLAGS: -L . -lntgcalls -Wl,-rpath=./
import "C"
import (
	"fmt"
	"main/ntgcalls"
	"os"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
	dotenv "github.com/joho/godotenv"
)

var caller *ntgcalls.Client
var client *tg.Client

func main() {
	dotenv.Load()
	caller = ntgcalls.NTgCalls()
	defer caller.Free()
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:    10029733,
		AppHash:  "d0d81009d46e774f78c0e0e622f5fa21",
		Session:  "session",
		LogLevel: tg.LogInfo,
	})
	client.Start()

	bot, _ := tg.NewClient(tg.ClientConfig{
		AppID:    10029733,
		AppHash:  "d0d81009d46e774f78c0e0e622f5fa21",
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
			caller.SetStreamSources(m.ChatID(), ntgcalls.PlaybackStream, ntgcalls.MediaDescription{
				Microphone: &ntgcalls.AudioDescription{
					MediaSource:  ntgcalls.MediaSourceFile,
					SampleRate:   128000,
					ChannelCount: 2,
					Input:        convertedFile,
				},
			})
			return nil
		}

		m.Reply("Error playing file: " + err.Error())
		return nil
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

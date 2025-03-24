package modules

import (
	"main/config"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func StartHandler(m *tg.NewMessage) error {
	message := "VCPlayBot is Active!"

	if config.StartImageUrl != "" {
		_, err := m.ReplyMedia(config.StartImageUrl, tg.MediaOptions{
			Caption: message,
		})
		return err
	} else {
		_, err := m.Reply(message)
		return err
	}
}

package modules

import (
	tg "github.com/amarnathcjd/gogram/telegram"
	"config"
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
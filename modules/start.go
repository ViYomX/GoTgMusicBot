package modules

import (
	"main/config"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func StartHandler(m *tg.NewMessage) error {
    message := `Hello 👋, I'm your Edit Guardian Bot, 
here to maintain a secure environment for our discussions.

🚫 Edited Message Deletion: I'll remove edited messages 
to maintain transparency.

📣 Notifications: You'll be informed each time a message is deleted.`

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

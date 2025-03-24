package modules

import (
	"fmt"

	"main/config"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func StartHandler(m *tg.NewMessage) error {
	user, _ := m.GetSender()
	mention := fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, user.ID, user.FirstName)

	message := fmt.Sprintf(`Hello %s 👋, I'm your Edit Guardian Bot, 
here to maintain a secure environment for our discussions.

<b>🚫 Edited Message Deletion: I'll remove edited messages 
to maintain transparency.</b>

📣 Notifications: You'll be informed each time a message is deleted.`, mention)

	if config.StartImageUrl != "" {
		_, err := m.ReplyMedia(config.StartImageUrl, tg.MediaOptions{
			Caption:   message,
			ParseMode: "HTML",
		})
		return err
	} else {
		_, err := m.Reply(message, tg.SendOptions{
			ParseMode: "HTML",
		})
		return err
	}
}

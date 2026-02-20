package bot

import (
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow/types/events"
)

func HandleEvent(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		text := v.Message.GetConversation()
		if text == "" {
			return
		}

		fmt.Printf("Message from %s: %s\n", v.Info.Sender.User, text)

		if !v.Info.IsFromMe {
			return
		}

		trimmed := strings.TrimSpace(text)
		if !strings.HasPrefix(trimmed, "*&") {
			return
		}

		keyword := strings.TrimSpace(strings.TrimPrefix(trimmed, "*&"))
		if keyword != "" {
			fmt.Printf("Command detected: %s\n", keyword)
		}
	}
}

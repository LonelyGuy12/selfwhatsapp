package bot

import (
	"context"
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

var Client *whatsmeow.Client

func HandleEvent(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		text := v.Message.GetConversation()
		if text == "" {
			return
		}

		fmt.Printf("Message from %s: %s\n", v.Info.Sender.User, text)

		if !v.Info.IsFromMe {
			if strings.EqualFold(strings.TrimSpace(text), "hi") {
				Client.SendMessage(context.Background(), v.Info.Chat, &waE2E.Message{
					Conversation: proto.String("hewwo senpai! :3"),
				})
			}
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

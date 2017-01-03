package main

import (
	"maunium.net/go/maubot"
	"strings"
)

func MessageHandler(message maubot.Message, mention, dm bool) {
	if strings.Contains(message.Text(), "ping") {
		message.Reply("pong")
	}
}

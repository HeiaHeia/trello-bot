package main

import (
	"maunium.net/go/maubot"
	"strings"
)

func MessageHandler(message maubot.Message) {
	if strings.Contains(message.Text(), "ping") {
		message.Reply("pong")
	}
}

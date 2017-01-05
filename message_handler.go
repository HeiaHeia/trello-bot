package main

import (
	"maunium.net/go/maubot"
	"regexp"
	"strconv"
	"strings"
	"trello-bot/trello"
)

func MessageHandler(message maubot.Message, mention, dm bool) {

	if strings.Contains(message.Text(), "ping") {
		message.Reply("pong")
	}

	if strings.HasPrefix(strings.ToLower(message.Text()), "report") {
		if strings.ToLower(message.Text()) == "report" {
			message.Reply("Usage: `report board:\"Board name\" days:\"14\" lists\"List, Another list`\"")
		} else {
			r := regexp.MustCompile("board:\"(.*)\" days:\"(.\\d)\" lists:\"(.*)\"")
			result := r.FindStringSubmatch(message.Text())
			board := result[1]
			days, err := strconv.Atoi(result[2])
			if err != nil {
				message.Reply("Days must be an integer")
				return
			}
			lists := strings.Split(result[3], ", ")
			conf := trello.ReportConfig{Board: board, Days: -days, Lists: lists}
			report, err := trello.GenerateReport(conf)
			if err != nil {
				message.Reply(err.Error())
				return
			}
			message.Reply(report)
		}
	}
}

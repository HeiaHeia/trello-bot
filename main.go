package main

import (
	"fmt"
	"github.com/aodin/date"
	"net/http"
	"trello-bot/slack"
	"trello-bot/trello"
)

var globalConfig Config

func main() {

	globalConfig = LoadConfig("config.json")

	trello.Authenticate(trello.TrelloConfig{Key: globalConfig.TrelloKey, Token: globalConfig.TrelloToken, User: globalConfig.TrelloUser, ActionHandler: ActionHandler})

	slack.Authenticate(globalConfig.SlackToken)

	go trello.SetupWebhooks(globalConfig.ListenURL)
	http.HandleFunc("/trello_webhook", trello.WebhookHandler)
	uid := RandomizeUID()
	err := slack.TryMessageChannelName(globalConfig.InfoChannel, fmt.Sprint("Trellobot report URL is ", globalConfig.ListenURL+"/"+uid))
	fmt.Println("Error sending report URL:", err)
	http.HandleFunc("/"+uid, reportHandler)
	listen := fmt.Sprintf(":%v", globalConfig.Port)
	fmt.Println("Starting server...")
	http.ListenAndServe(listen, nil)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Loading Report")
	report, err := trello.CompletedReport(date.NewRange(date.Today().AddDays(globalConfig.ReportDays), date.Today()), globalConfig.ReportLists)
	if err != nil {
		fmt.Fprintln(w, "Error loading report")
		return
	}
	fmt.Fprintln(w, report)
}

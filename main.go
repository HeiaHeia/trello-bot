package main

import (
	"fmt"
	"github.com/aodin/date"
	"net/http"
	"time"
	"trello-bot/slack"
	"trello-bot/trello"
)

var globalConfig Config

func main() {

	globalConfig = LoadConfig("config.json")

	trello.Authenticate(trello.TrelloConfig{Key: globalConfig.TrelloKey, Token: globalConfig.TrelloToken, User: globalConfig.TrelloUser, ActionHandler: ActionHandler})
	slack.Authenticate(globalConfig.SlackToken)

	for i := range globalConfig.BoardConfigs {
		boardName := globalConfig.BoardConfigs[i].BoardName
		go setupWebhook(boardName)
	}
	http.HandleFunc("/trello_webhook", trello.WebhookHandler)
	uid := RandomizeUID()
	err := slack.TryMessageChannelName(globalConfig.InfoChannel, fmt.Sprint("Trellobot report URL is ", globalConfig.ListenURL+"/"+uid))
	if err != nil {
		fmt.Println("Error sending report URL:", err)
	}
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

func setupWebhook(boardName string) {

	fmt.Println("Waiting for the server to start up before creating webhooks")
	time.Sleep(2 * time.Second)
	trello.SetupWebhook(boardName, globalConfig.ListenURL)
}

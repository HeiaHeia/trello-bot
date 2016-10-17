package main

import (
	"fmt"
	"net/http"
	"trello-bot/slack"
	"trello-bot/trello"
)

func main() {

	config := LoadConfig("config.json")

	trello.Authenticate(trello.TrelloConfig{
		Key:              config.TrelloKey,
		Token:            config.TrelloToken,
		User:             config.TrelloUser,
		Board:            config.BoardName,
		StartingListName: config.StartListName,
		FinishedListName: config.FinishedListName})

	slack.Authenticate(config.SlackToken)

	go trello.SetupWebhooks(config.ListenURL)
	http.HandleFunc("/trello_webhook", trello.WebhookHandler)
	listen := fmt.Sprintf(":%v", config.Port)
	fmt.Println("Starting server...")
	http.ListenAndServe(listen, nil)
}

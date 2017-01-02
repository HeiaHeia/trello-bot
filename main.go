package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"trello-bot/slack"
	"trello-bot/trello"
)

var globalConfig Config

func main() {

	var err error
	globalConfig, err = LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	trello.Setup(trello.TrelloConfig{Key: globalConfig.TrelloKey, Token: globalConfig.TrelloToken, User: globalConfig.TrelloUser, ActionHandler: ActionHandler})
	go slack.Start(globalConfig.SlackToken, MessageHandler)

	for i := range globalConfig.BoardConfigs {
		boardName := globalConfig.BoardConfigs[i].BoardName
		go setupWebhook(boardName)
	}
	http.HandleFunc("/trello_webhook", trello.WebhookHandler)
	listen := fmt.Sprintf(":%v", globalConfig.Port)
	fmt.Println("Starting server...")
	http.ListenAndServe(listen, nil)
}

func setupWebhook(boardName string) {

	fmt.Println("Waiting for the server to start up before creating webhooks")
	time.Sleep(2 * time.Second)
	trello.SetupWebhook(boardName, globalConfig.ListenURL)
}

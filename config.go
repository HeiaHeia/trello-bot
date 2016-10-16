package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TrelloUser       string `json:"trello-user"`
	TrelloKey        string `json:"trello-key"`
	TrelloToken      string `json:"trello-token"`
	BoardName        string `json:"board-name"`
	StartListName    string `json:"start-list-name"`
	FinishedListName string `json:"finished-list-name"`
	SlackToken       string `json:"slack-token"`
	Port             string `json:"port"`
}

func LoadConfig(filename string) Config {

	var config Config

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	json.Unmarshal(file, &config)

	return config
}

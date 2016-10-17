package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TrelloUser       string `json:"trello_user"`
	TrelloKey        string `json:"trello_key"`
	TrelloToken      string `json:"trello_token"`
	BoardName        string `json:"board_name"`
	StartListName    string `json:"start_list_name"`
	FinishedListName string `json:"finished_list_name"`
	SlackToken       string `json:"slack_token"`
	ListenURL        string `json:"listen_url"`
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

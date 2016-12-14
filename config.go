package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TrelloUser  string `json:"trello_user"`
	TrelloKey   string `json:"trello_key"`
	TrelloToken string `json:"trello_token"`
	SlackToken  string `json:"slack_token"`
	ListenURL   string `json:"listen_url"`
	Port        string `json:"port"`
}

type BoardConfig struct {
	BoardName         string       `json:"board_name"`
	NotifyChannelName string       `json:"notify_channel_name"`
	ListConfigs       []ListConfig `json:"list_configs"`
}

type ListConfig struct {
	ListName        string `json:"list_name"`
	OnAction        string `json:"on_action"`
	MessageTemplate string `json:"message_template"`
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

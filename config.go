package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	TrelloUser   string        `json:"trello_user"`
	TrelloKey    string        `json:"trello_key"`
	TrelloToken  string        `json:"trello_token"`
	SlackToken   string        `json:"slack_token"`
	ListenURL    string        `json:"listen_url"`
	Port         string        `json:"port"`
	BoardConfigs []BoardConfig `json:"board_configs"`
}

type BoardConfig struct {
	BoardName         string       `json:"board_name"`
	NotifyChannelName string       `json:"notify_channel_name"`
	OnAction          string       `json:"on_action"`
	MessageTemplate   string       `json:"message_template"`
	ListConfigs       []ListConfig `json:"list_configs"`
}

type ListConfig struct {
	ListName        string `json:"list_name"`
	OnAction        string `json:"on_action"`
	MessageTemplate string `json:"message_template"`
}

const (
	ActionMovedTo   string = "moved_to"
	ActionMovedFrom string = "moved_from"
)

func LoadConfig() (Config, error) {

	filename, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	var config Config

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file: %v\n", err)
	}

	json.Unmarshal(file, &config)

	return config, nil
}

func getConfigPath() (string, error) {

	envPath := os.Getenv("TRELLOBOT_CONFIG_PATH")
	defaultPath := "/etc/trellobot/trellobot.conf"

	for _, path := range []string{envPath, defaultPath} {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", errors.New("Config file location not specified")
}

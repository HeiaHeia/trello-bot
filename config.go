package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	TrelloUser   string        `yaml:"trello_user"`
	TrelloKey    string        `yaml:"trello_key"`
	TrelloToken  string        `yaml:"trello_token"`
	SlackToken   string        `yaml:"slack_token"`
	ListenURL    string        `yaml:"listen_url"`
	Port         string        `yaml:"port"`
	BoardConfigs []BoardConfig `yaml:"board_configs"`
}

type BoardConfig struct {
	BoardName         string       `yaml:"board_name"`
	NotifyChannelName string       `yaml:"notify_channel_name"`
	OnAction          string       `yaml:"on_action"`
	MessageTemplate   string       `yaml:"message_template"`
	ListConfigs       []ListConfig `yaml:"list_configs"`
}

type ListConfig struct {
	ListName        string `yaml:"list_name"`
	OnAction        string `yaml:"on_action"`
	MessageTemplate string `yaml:"message_template"`
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

	yaml.Unmarshal(file, &config)

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

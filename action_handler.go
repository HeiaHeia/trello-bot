package main

import (
	"github.com/joonasmyhrberg/go-trello"
	"strings"
	"trello-bot/slack"
)

const (
	TypeUpdateCard      string = "updateCard"
	TypeMoveCardToBoard string = "moveCardToBoard"
)

const (
	TemplateCardName string = "[card_name]"
	TemplateCardLink string = "[card_link]"
)

func ActionHandler(action trello.Action) {
	for _, boardConfig := range globalConfig.BoardConfigs {
		if action.Data.Board.Name != boardConfig.BoardName {
			continue
		}
		checkBoardAction(boardConfig, action)
		for _, listConfig := range boardConfig.ListConfigs {
			checkListAction(listConfig, boardConfig, action)
		}
	}
}

func checkListAction(listConfig ListConfig, boardConfig BoardConfig, action trello.Action) {
	if listConfig.OnAction == ActionMovedTo {
		if action.Data.ListAfter.Name == listConfig.ListName {
			sendMessage(listConfig.MessageTemplate, boardConfig.NotifyChannelName, action)
		}
	} else {
		if action.Data.ListBefore.Name == listConfig.ListName {
			sendMessage(listConfig.MessageTemplate, boardConfig.NotifyChannelName, action)
		}
	}
}

func checkBoardAction(boardConfig BoardConfig, action trello.Action) {
	if action.Type == TypeMoveCardToBoard &&
		boardConfig.OnAction == ActionMovedTo &&
		boardConfig.BoardName == action.Data.Board.Name {
		sendMessage(boardConfig.MessageTemplate, boardConfig.NotifyChannelName, action)
	}
}

func sendMessage(template, channel string, action trello.Action) {
	message := parseMessage(template, action.Data.Card.Name, "https://trello.com/c/"+action.Data.Card.ShortLink)
	slack.TryMessageChannelName(channel, message)
}

func parseMessage(template, name, link string) (message string) {
	message = strings.Replace(template, TemplateCardName, name, -1)
	message = strings.Replace(message, TemplateCardLink, link, -1)
	return
}

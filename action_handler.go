package main

import (
	"strings"
	"trello-bot/slack"
	trelloHelper "trello-bot/trello"

	"github.com/joonasmyhrberg/go-trello"
)

const (
	TypeUpdateCard      string = "updateCard"
	TypeMoveCardToBoard string = "moveCardToBoard"
)

const (
	TemplateCardName   string = "[card_name]"
	TemplateCardLink   string = "[card_link]"
	TemplateCardLabels string = "[card_labels]"
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
	message := parseMessage(action.Data.Card.Id, template, action.Data.Card.Name, "https://trello.com/c/"+action.Data.Card.ShortLink)
	slack.TryMessageChannelName(channel, message)
}

func parseMessage(cardID, template, name, link string) (message string) {
	message = strings.Replace(template, TemplateCardName, name, -1)
	message = strings.Replace(message, TemplateCardLink, link, -1)
	labels := trelloHelper.GetCardLabels(cardID)
	if len(labels) > 0 {
		labelsString := strings.Join(labels, ", ") + ":"
		message = strings.Replace(message, TemplateCardLabels, labelsString, -1)
	} else {
		message = strings.Replace(message, TemplateCardLabels, "", -1)
	}
	return
}

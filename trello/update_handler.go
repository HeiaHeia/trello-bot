package trello

import (
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
	"trello-bot/slack"
)

func handleAction(action trello.Action) {

	switch action.Type {
	case "updateCard":
		cardUpdated(action)
	case "moveCardToBoard":
		cardAccepted(action)
	}
}

func cardUpdated(action trello.Action) {

	newList := action.Data.ListAfter
	switch newList.Name {
	case trelloConfig.StartingListName:
		cardStarted(action.Data.Card.Id)
	case trelloConfig.FinishedListName:
		cardFinished(action.Data.Card.Id)
	}
}

func cardStarted(cardID string) {

	card, err := trelloClient.Card(cardID)
	if err != nil {
		return
	}

	sendMessageForCard(card, fmt.Sprintf("%s is now in development", card.Name))
}

func cardFinished(cardID string) {

	card, err := trelloClient.Card(cardID)
	if err != nil {
		return
	}

	sendMessageForCard(card, fmt.Sprintf("%s is now finished", card.Name))
}

func cardAccepted(action trello.Action) {

	card, err := trelloClient.Card(action.Data.Card.Id)
	if err != nil {
		return
	}

	sendMessageForCard(card, fmt.Sprintf("%s has been added to %s", card.Name, trelloConfig.Board))
}

func sendMessageForCard(card *trello.Card, message string) {

	comments, err := commentsForCard(card)
	if err != nil {
		return
	}
	for _, comment := range comments {
		_ = slack.TryMessageUsername(comment, message)
	}
}

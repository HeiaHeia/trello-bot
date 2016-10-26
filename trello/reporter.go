package trello

import (
	"fmt"
	"github.com/aodin/date"
	"github.com/joonasmyhrberg/go-trello"
	"strings"
	"time"
)

func CompletedReport(dateRange date.Range, listNames []string) (report string, err error) {

	board, err := getMainBoard()
	if err != nil {
		return "", err
	}

	actions, err := board.Actions(trello.NewArgument("filter", "updateCard:idList"), trello.NewArgument("limit", "500"))
	if err != nil {
		return "", err
	}

	var completedActions []trello.Action

	for _, action := range actions {
		time, err := time.Parse(time.RFC3339Nano, action.Date)
		if err != nil {
			return "", err
		}
		inRange := date.FromTime(time).Within(dateRange)
		finished := listInLists(action.Data.ListAfter.Name, listNames)
		if inRange && finished {
			completedActions = append(completedActions, action)
		}
	}

	completedCards, err := cardsInActions(completedActions)
	completedCards = removeDuplicates(completedCards)
	labelCounts := labelCounts(completedCards)

	report += fmt.Sprintf("%v cards completed", len(completedCards))
	var countReports []string
	for label, count := range labelCounts {
		desc := fmt.Sprintf("%s: %v", label, count)
		countReports = append(countReports, desc)
	}
	report += " (" + strings.Join(countReports, ", ") + ")\n  "
	var cardReports []string
	for _, card := range completedCards {
		desc := fmt.Sprintf("%s %s", card.Name, card.ShortUrl)
		cardReports = append(cardReports, desc)
	}
	report += strings.Join(cardReports, "\n  ")

	return
}

func removeDuplicates(cards []*trello.Card) []*trello.Card {

	cardMap := make(map[string]*trello.Card)

	for _, card := range cards {
		cardMap[card.Id] = card
	}

	var deduplicated []*trello.Card

	for _, card := range cardMap {
		deduplicated = append(deduplicated, card)
	}

	return deduplicated
}

func cardsInActions(actions []trello.Action) (cards []*trello.Card, err error) {

	for _, action := range actions {
		card, err := trelloClient.Card(action.Data.Card.Id)
		if err != nil {
			return []*trello.Card{}, err
		}
		cards = append(cards, card)
	}

	return
}

func labelCounts(cards []*trello.Card) map[string]int {

	counts := make(map[string]int)

	for _, card := range cards {
		for _, label := range card.Labels {
			counts[label.Name] += 1
		}
	}

	return counts
}

func listInLists(listName string, listNames []string) bool {

	for _, l := range listNames {
		if l == listName {
			return true
		}
	}

	return false
}

package trello

import (
	"fmt"
	"github.com/aodin/date"
	"github.com/joonasmyhrberg/go-trello"
	"strings"
	"time"
)

type ReportConfig struct {
	Board string
	Days  int
	Lists []string
}

func GenerateReport(reportConfig ReportConfig) (report string, err error) {

	dateRange := date.NewRange(date.Today().AddDays(reportConfig.Days), date.Today())

	actions, err := actionsFromBoard(reportConfig.Board, trelloUser)
	if err != nil {
		return "", err
	}

	filteredActions, err := filterActions(actions, dateRange, reportConfig.Lists)
	if err != nil {
		return "", err
	}

	cards, err := cardsInActions(filteredActions)
	if err != nil {
		return "", err
	}
	cards = removeDuplicates(cards)
	labelCounts := labelCounts(cards)

	report += fmt.Sprintf("%v cards completed", len(cards))
	var countReports []string
	for label, count := range labelCounts {
		desc := fmt.Sprintf("%s: %v", label, count)
		countReports = append(countReports, desc)
	}
	report += " (" + strings.Join(countReports, ", ") + ")\n  "
	var cardReports []string
	for _, card := range cards {
		desc := fmt.Sprintf("%s %s", card.Name, card.ShortUrl)
		cardReports = append(cardReports, desc)
	}
	report += strings.Join(cardReports, "\n  ")

	return
}

func actionsFromBoard(boardName string, user *trello.Member) ([]trello.Action, error) {

	board, err := getBoard(boardName, user)
	if err != nil {
		return []trello.Action{}, err
	}

	actionsFilter := trello.NewArgument("filter", "updateCard:idList")
	limit := trello.NewArgument("limit", "500")
	actions, err := board.Actions(actionsFilter, limit)
	if err != nil {
		return []trello.Action{}, err
	}

	return actions, nil
}

func filterActions(actions []trello.Action, dateRange date.Range, lists []string) ([]trello.Action, error) {

	var filteredActions []trello.Action

	for _, action := range actions {
		time, err := time.Parse(time.RFC3339Nano, action.Date)
		if err != nil {
			return []trello.Action{}, err
		}
		inRange := date.FromTime(time).Within(dateRange)
		inLists := listInLists(action.Data.ListAfter.Name, lists)
		if inRange && inLists {
			filteredActions = append(filteredActions, action)
		}
	}

	return filteredActions, nil
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

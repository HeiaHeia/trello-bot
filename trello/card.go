package trello

import (
//"github.com/joonasmyhrberg/go-trello"
)

func GetCardLabels(cardID string) (labels []string) {

	card, err := trelloClient.Card(cardID)
	if err != nil {
		return
	}

	for _, label := range card.Labels {
		labels = append(labels, label.Name)
	}

	return
}

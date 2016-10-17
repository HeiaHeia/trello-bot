package trello

import (
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
)

func commentsForCard(card *trello.Card) (comments []string, err error) {

	actions, err := card.Actions()
	if err != nil {
		fmt.Println("Error retrieving card actions: ", err)
		return []string{}, err
	}

	for _, action := range actions {
		comment := action.Data.Text
		if len(comment) > 0 {
			comments = append(comments, comment)
		}
	}

	return
}

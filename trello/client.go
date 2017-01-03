package trello

import (
	"github.com/joonasmyhrberg/go-trello"
)

var trelloClient *trello.Client
var trelloUser *trello.Member

func Setup(config TrelloConfig) error {

	trelloConfig = config

	trello, err := trello.NewAuthClient(config.Key, &config.Token)
	if err != nil {
		return err
	}
	trelloClient = trello

	user, err := trello.Member(config.User)
	if err != nil {
		return err
	}
	trelloUser = user

	return nil
}

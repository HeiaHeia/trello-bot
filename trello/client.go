package trello

import (
	"github.com/joonasmyhrberg/go-trello"
)

var trelloClient *trello.Client

func Authenticate(config TrelloConfig) error {

	trelloConfig = config

	trello, err := trello.NewAuthClient(config.Key, &config.Token)
	trelloClient = trello

	return err
}

package trello

import (
	"github.com/joonasmyhrberg/go-trello"
)

type TrelloConfig struct {
	Key              string
	Token            string
	User             string
	Board            string
	StartingListName string
	FinishedListName string
	NotifyChannel    string
}

var trelloClient *trello.Client
var trelloConfig TrelloConfig

func Authenticate(config TrelloConfig) error {

	trelloConfig = config

	trello, err := trello.NewAuthClient(config.Key, &config.Token)
	trelloClient = trello

	return err
}

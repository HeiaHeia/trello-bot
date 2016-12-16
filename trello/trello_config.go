package trello

import (
	"github.com/joonasmyhrberg/go-trello"
)

type TrelloConfig struct {
	Key           string
	Token         string
	User          string
	ActionHandler func(action trello.Action)
}

var trelloConfig TrelloConfig

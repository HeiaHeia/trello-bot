package trello

import (
	"errors"
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
)

func getMainBoard() (board trello.Board, err error) {

	user, err := trelloClient.Member(trelloConfig.User)
	if err != nil {
		return trello.Board{}, err
	}

	return getBoard(trelloConfig.Board, user)
}

func getBoard(name string, user *trello.Member) (board trello.Board, err error) {

	boards, err := user.Boards()
	if err != nil {
		return trello.Board{}, err
	}

	for _, board := range boards {
		if board.Name == name {
			return board, nil
		}
	}

	return trello.Board{}, errors.New(fmt.Sprintf("no board with name %s", name))
}

package slack

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
)

func getUserWithUsername(username string) (user slack.User, err error) {

	users, err := slackClient.GetUsers()
	if err != nil {
		return slack.User{}, err
	}

	for _, user := range users {
		if user.Name == username {
			return user, nil
		}
	}

	return slack.User{}, errors.New(fmt.Sprintf("No user with username %s", username))
}

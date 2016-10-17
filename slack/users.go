package slack

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"strings"
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

func parseUsername(name string) (username string, err error) {

	username = strings.ToLower(name)

	if strings.Contains(username, " ") {
		return "", errors.New(fmt.Sprintf("%s is not a Slack username", name))
	}

	if strings.Contains(name, "@") {
		username = strings.Replace(username, "@", "", 1)
	}

	return
}

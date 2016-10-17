package slack

import (
	"github.com/nlopes/slack"
)

func DmUpdate(name string, message string) {

	username, err := parseUsername(name)
	if err != nil {
		return
	}

	user, err := getUserWithUsername(username)
	if err != nil {
		return
	}

	err = sendDm(user, message)
}

func sendDm(user slack.User, message string) error {

	_, _, channel, err := slackClient.OpenIMChannel(user.ID)
	if err != nil {
		return err
	}

	_, _, err = slackClient.PostMessage(channel, message, slack.PostMessageParameters{AsUser: true})

	return err
}

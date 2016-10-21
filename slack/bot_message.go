package slack

import (
	"github.com/nlopes/slack"
)

func TryMessageChannelName(channelName, message string) error {

	channel, err := getChannel(channelName)
	if err != nil {
		return err
	}

	err = messageChannel(channel, message)

	return err
}

func TryMessageUsername(username, message string) error {

	parsedName, err := parseUsername(username)
	if err != nil {
		return err
	}
	user, err := getUserWithUsername(parsedName)
	if err != nil {
		return err
	}

	err = messageUser(user, message)

	return err
}

func messageChannel(channel slack.Channel, message string) error {

	params := slack.PostMessageParameters{AsUser: true}
	_, _, err := slackClient.PostMessage(channel.ID, message, params)

	return err
}

func messageUser(user slack.User, message string) error {

	_, _, channel, err := slackClient.OpenIMChannel(user.ID)
	if err != nil {
		return err
	}

	_, _, err = slackClient.PostMessage(channel, message, slack.PostMessageParameters{AsUser: true})

	return err
}

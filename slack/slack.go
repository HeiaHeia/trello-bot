package slack

import (
	slackAPI "github.com/nlopes/slack"
	"maunium.net/go/maubot"
	"maunium.net/go/maubot/slack"
)

var bot maubot.Maubot
var slackClient slackAPI.RTM
var slackUID string

func Start(token string, messageHandler func(message maubot.Message)) error {

	bot = maubot.New()
	slackBot, err := slack.New(token)
	slackUID = slackBot.UID()
	slackRTM, ok := slackBot.Underlying().(slackAPI.RTM)
	if ok {
		slackClient = slackRTM
	}
	if err != nil {
		return err
	}
	err = slackBot.Connect()
	if err != nil {
		return err
	}
	bot.Add(slackBot)
	for message := range bot.Messages() {
		messageHandler(message)
	}

	return nil
}

func TryMessageChannelName(channelName, message string) error {

	channel, err := getChannel(channelName)
	if err != nil {
		return err
	}

	messageChannel(channel, message)

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

func messageChannel(channel slackAPI.Channel, message string) {

	outgoingMessage := maubot.OutgoingMessage{Text: message, RoomID: channel.ID, PlatformID: slackUID}
	bot.SendMessage(outgoingMessage)
}

func messageUser(user slackAPI.User, message string) error {

	_, _, channel, err := slackClient.OpenIMChannel(user.ID)
	if err != nil {
		return err
	}

	outgoingMessage := maubot.OutgoingMessage{Text: message, RoomID: channel, PlatformID: slackUID}
	bot.SendMessage(outgoingMessage)

	return nil
}

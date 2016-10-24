package slack

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
)

func getChannel(name string) (channel slack.Channel, err error) {

	channels, err := slackClient.GetChannels(true)
	if err != nil {
		return
	}

	for _, channel := range channels {
		if channel.Name == name {
			return channel, nil
		}
	}

	return slack.Channel{}, errors.New(fmt.Sprintf("No channel named %s", name))
}

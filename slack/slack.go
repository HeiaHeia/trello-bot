package slack

import (
	"github.com/nlopes/slack"
)

var slackClient *slack.Client

func Setup(token string) {

	slackClient = slack.New(token)
}

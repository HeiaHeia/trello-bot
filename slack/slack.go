package slack

import (
	"github.com/nlopes/slack"
)

var slackClient *slack.Client

func Authenticate(token string) {

	slackClient = slack.New(token)
}

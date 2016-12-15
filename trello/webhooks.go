package trello

import (
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
	"net/url"
	"path"
	"strings"
)

func SetupWebhook(boardName, listenURL string) {

	user, err := trelloClient.Member(trelloConfig.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	targetBoard, err := getBoard(boardName, user)
	if err != nil {
		fmt.Println(err)
		return
	}
	webhooks, err := trelloClient.Webhooks(trelloConfig.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, webhook := range webhooks {
		if webhook.IDModel == targetBoard.Id && strings.Contains(webhook.CallbackURL, listenURL) {
			fmt.Println("Webhook with correct ID and URL already exists")
			return
		}
	}

	fmt.Println("Creating a new webhook")
	err = createWebhook(targetBoard.Id, listenURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func createWebhook(modelID, baseURL string) error {

	webhookURL, err := makeWebhookURL(baseURL)
	webhook := trello.Webhook{IDModel: modelID, CallbackURL: webhookURL, Description: "Trellobot Monitoring Webhook"}
	_, err = trelloClient.CreateWebhook(webhook)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func makeWebhookURL(baseURL string) (string, error) {

	URL, err := url.Parse(baseURL)
	URL.Path = path.Join(URL.Path, "trello_webhook")

	return URL.String(), err
}

package trello

import (
	"errors"
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
	"net/url"
	"path"
	"strings"
	"time"
)

func SetupWebhooks(listenURL string) {

	targetBoard, err := findBoard(trelloConfig.Board)
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

	fmt.Println("Waiting 2 seconds for the server to start up")
	time.Sleep(2 * time.Second)

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

func findBoard(name string) (trello.Board, error) {

	user, err := trelloClient.Member(trelloConfig.User)
	if err != nil {
		fmt.Println(err)
		return trello.Board{}, err
	}
	boards, err := user.Boards()
	if err != nil {
		fmt.Println(err)
		return trello.Board{}, err
	}
	for _, board := range boards {
		if board.Name == name {
			return board, nil
		}
	}
	return trello.Board{}, errors.New("No boards with given name")
}

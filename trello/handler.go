package trello

import (
	"encoding/json"
	"fmt"
	"github.com/joonasmyhrberg/go-trello"
	"io/ioutil"
	"net/http"
)

type WebhookResponse struct {
	Action trello.Action `json:"action"`
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
		handlePing(w)
	case "POST":
		response, err := handleUpdate(w, r)
		if err != nil {
			fmt.Println("Error handling update: ", err)
		}
		trelloConfig.ActionHandler(response.Action)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handlePing(w http.ResponseWriter) {
	fmt.Println("Received webhook creation ping, responding with 200")
	w.WriteHeader(http.StatusOK)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) (webhookResponse WebhookResponse, err error) {

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &webhookResponse)
	if err != nil {
		fmt.Println("Failed to decode action JSON with error: ", err.Error())
		return WebhookResponse{}, err
	}
	return
}

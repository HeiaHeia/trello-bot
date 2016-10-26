# Trellobot

Automatic Slack notifications for Trello-based development workflows.

- Add a comment with your Slack username to a card and get notifications when the card has been:
  - Approved for development (moved to board)
  - Started (moved to in progress list)
  - Finished (moved to finished list)
- Send notifications to a channel when any card has been moved to finished list
- Download a report of cards finished in a time frame

## Installation

The recommended installation method is with Docker (Docker quick install: `$ curl -sSL https://get.docker.com/ | sh`). After downloading the project:

1. Fill the values in `config.json`
2. Run `$ make run-detached`

## Configuration

- `"trello_user" `Your Trello username. If this is a shared bot, it is recommended to create a new account for it.
- `"trello_key"` Trello API key. Can be retrieved from <https://trello.com/app-key>.
- `"trello_token"` Trello API token. You can use the test token from the site above.
- `"board_name"` The name of the board that the bot will observe.
- `"start_list_name"` Cards moved to this list are considered started.
- `"finished_list_name"` List for finished cards.
- `"notify_channel" (optional) `Name of the channel where notifications should be sent to. If this is empty, notifications will be disabled.
- `"info_channel"` (optional) Name of the channel where the bot should send info messages (e.g. report URL).
- `"report_days"` Number of days backwards the report should span. Should be a negative integer.
- `"report_lists"` The lists that should be included in the report. If you have multiple finished states (ready, shipped) all of those lists should be added here. Otherwise it should be the same as `finished_list_name`. Should be an array of strings (list names).
- `"slack_token"` Slack API token.
- `"listen_url"` The URL of the server. This will be used for creating the webhooks.
- `"port"` Port where the server should listen. Should almost always be 80. For TLS it is recommended to add a reverse proxy (like Nginx) to forward the requests with TLS stripped.

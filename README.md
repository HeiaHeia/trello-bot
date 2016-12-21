# Trellobot

Automatic Slack notifications for Trello-based development workflows. Slack notifications can be triggered when a card was moved to/from board or list.

## Installation

The recommended installation method is with Docker (Docker quick install: `$ curl -sSL https://get.docker.com/ | sh`). After downloading the project:

1. Copy the example config: `$ cp config.example.yaml config.yaml`
2. Fill the values in the config.
3. Run `$ docker run -d -p 80:80 $PWD/config.yaml:/etc/trellobot/trellobot.conf --name trellobot joonasmyhrberg/trello-bot`

## Configuration

Example configuration file can be found from [config.example.yaml].

### Application Configuration
- `trello_user` Your Trello username. If this is a shared bot, it is recommended to create a new account for it.
- `trello_key` Trello API key. Can be retrieved from <https://trello.com/app-key>.
- `trello_token` Trello API token. You can use the test token from the site above.
- `slack_token` Slack API token.
- `listen_url` The URL of the server. This will be used for creating the webhooks.
- `port` Port where the server should listen. Should almost always be 80. For TLS it is recommended to add a reverse proxy (like Nginx) to forward the requests with TLS stripped.

### Notification Configuration
- `board_configs` Configurations for all the boards that should be monitored.
  - `board_name` Name of the board.
  - `notify_channel` Name of the Slack channels where notifications should be sent to.
  - `list_configs` Configurations for all the lists that should be monitored.
    - `list_name` Name of the list.
    - `on_action` The action that should trigger the notification. Can be `moved_to` or `moved_from`.
    - `message_template` The template that will be used to render the notification. Possible placeholders: `[card_name]` and `[card_link]`

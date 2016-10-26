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

1. Fill the values in `confif.json`
2. Run `$ make run-detached`

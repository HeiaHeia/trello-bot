FROM golang:latest

ADD . /go/src/trello-bot
WORKDIR /go/src/trello-bot

ENV GOPATH /go
RUN go get

ENTRYPOINT /go/bin/trello-bot

NAME=trellobot

build:
	docker build -t $(NAME) .

run: build
	docker run --rm -it -p 80:80 -v ~/config.yaml:/etc/trellobot/trellobot.conf --name $(NAME) $(NAME)

run-detached: build
	docker run -d -p 80:80 -v ~/config.yaml:/etc/trellobot/trellobot.conf --name $(NAME) $(NAME)

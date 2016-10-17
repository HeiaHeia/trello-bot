NAME=trellobot

build:
	docker build -t $(NAME) .

run: build
	docker run --rm -it -p 80:80 --name $(NAME) $(NAME)

run-detached: build
	docker run -d -p 80:80 --name $(NAME) $(NAME)

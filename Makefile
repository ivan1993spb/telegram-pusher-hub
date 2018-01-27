IMAGE_NAME = "ivan1993spb/telegram-pusher-hub"

go/build:
	@docker run --rm -v $(PWD):/go/src/github.com/ivan1993spb/telegram-pusher-hub -w /go/src/github.com/ivan1993spb/telegram-pusher-hub golang:1.9-alpine go build -x -v -o telegram-pusher-hub

docker/build:
	@docker build -t $(IMAGE_NAME) .

docker/push:
	@docker push $(IMAGE_NAME)

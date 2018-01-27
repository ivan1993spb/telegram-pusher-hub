FROM alpine:latest

RUN apk add --no-cache ca-certificates

ADD telegram-pusher-hub /telegram-pusher-hub

ENTRYPOINT [ "/telegram-pusher-hub" ]

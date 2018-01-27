FROM alpine:latest

ADD telegram-pusher-hub /telegram-pusher-hub

ENTRYPOINT [ "/telegram-pusher-hub" ]

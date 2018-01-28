package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

const CacheSize = 1024

const BotBuffer = 100

var (
	ChannelUsername string
	ListenAddress   string
	TelegramToken   string
	SendingTimeout  time.Duration
)

func init() {
	flag.StringVar(&ChannelUsername, "channel_username", "", "Channel username like @channel")
	flag.StringVar(&ListenAddress, "listen_address", ":8080", "Address to serve")
	flag.StringVar(&TelegramToken, "token", "", "Telegram token xxx:yyy")
	flag.DurationVar(&SendingTimeout, "sending_timeout", time.Second, "Sending message timeout")
	flag.Parse()
}

func main() {
	log := logrus.New()
	log.Infof("Starting pushing to channel %q", ChannelUsername)

	bot := &tgbotapi.BotAPI{
		Token: TelegramToken,
		Client: &http.Client{
			Timeout: SendingTimeout,
		},
		Buffer: BotBuffer,
	}

	sender := NewSender(bot, ChannelUsername)
	pusher := NewPusher(CacheSize, sender, log)

	n := negroni.Classic()
	r := mux.NewRouter()
	apiV1Router := mux.NewRouter().PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	apiV1Router.Path(URLPathPushMessage).Methods(http.MethodPost).Handler(NewPushMessageHandler(pusher, log))

	r.PathPrefix("/api/v1").Handler(negroni.New(
		negroni.Wrap(apiV1Router),
	))

	n.UseHandler(r)
	go pusher.Run()
	n.Run(ListenAddress)
}

package main

import (
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

const CacheSize = 1024

const BotBuffer = 100

func main() {
	channelUsername := os.Getenv("CHANNEL_USERNAME")
	token := os.Getenv("TOKEN")
	address := os.Getenv("ADDRESS")

	log := logrus.New()
	log.Info("Starting")

	bot := &tgbotapi.BotAPI{
		Token:  token,
		Client: &http.Client{},
		Buffer: BotBuffer,
	}

	sender := NewSender(bot, channelUsername)
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
	n.Run(address)
}

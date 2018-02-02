package main

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const URLPathPushMessage = "/message"

const (
	FieldText                  = "text"
	FieldDisableNotification   = "disable_notification"
	FieldDisableWebPagePreview = "disable_webpage_preview"
	FieldParseMode             = "parse_mode"
)

type errPushMessage string

func (e errPushMessage) Error() string {
	return "error on push message handler: " + string(e)
}

type PushMessageHandler struct {
	pusher PusherInterface
	log    *logrus.Logger
}

func NewPushMessageHandler(pusher PusherInterface, log *logrus.Logger) http.Handler {
	return &PushMessageHandler{
		pusher: pusher,
		log:    log,
	}
}

func (h *PushMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Create secret api token.

	text := r.PostFormValue(FieldText)
	if len(text) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	disableNotification, _ := strconv.ParseBool(r.PostFormValue(FieldDisableNotification))
	disableWebPagePreview, _ := strconv.ParseBool(r.PostFormValue(FieldDisableWebPagePreview))
	parseMode := GetParseMode(r.PostFormValue(FieldParseMode))

	err := h.pusher.Push(&Message{
		Text:                  text,
		DisableNotification:   disableNotification,
		DisableWebPagePreview: disableWebPagePreview,
		ParseMode:             parseMode,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
}

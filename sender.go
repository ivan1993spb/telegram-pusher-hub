package main

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Sender struct {
	bot             *tgbotapi.BotAPI
	channelUsername string
}

func NewSender(bot *tgbotapi.BotAPI, channelUsername string) *Sender {
	return &Sender{
		bot:             bot,
		channelUsername: channelUsername,
	}
}

func (s *Sender) Send(msg *Message) error {
	_, err := s.bot.Send(&tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChannelUsername:     s.channelUsername,
			DisableNotification: msg.DisableNotification,
		},
		Text: msg.Text,
		DisableWebPagePreview: msg.DisableWebPagePreview,
		ParseMode:             msg.ParseMode.String(),
	})

	if err != nil {
		return fmt.Errorf("cannot send message: %s", err)
	}

	return nil
}

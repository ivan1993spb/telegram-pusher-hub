package main

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Pusher struct {
	ch          chan *Message
	sender      SenderInterface
	log         *logrus.Logger
	pushTimeout time.Duration
}

func NewPusher(cacheSize uint, sender SenderInterface, log *logrus.Logger, pushTimeout time.Duration) *Pusher {
	return &Pusher{
		ch:          make(chan *Message, cacheSize),
		sender:      sender,
		log:         log,
		pushTimeout: pushTimeout,
	}
}

func (p *Pusher) Push(msg *Message) error {
	timer := time.NewTimer(p.pushTimeout)
	defer timer.Stop()

	select {
	case p.ch <- msg:
	case <-timer.C:
		return errors.New("message pushing time is out")
	}

	return nil
}

func (p *Pusher) Run() {
	for msg := range p.ch {
		if err := p.sender.Send(msg); err != nil {
			p.log.Errorf("Sending message error: %s", err)
		}
	}
}

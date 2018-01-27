package main

import "github.com/sirupsen/logrus"

type Pusher struct {
	ch     chan *Message
	sender SenderInterface
	log    *logrus.Logger
}

func NewPusher(cacheSize uint, sender SenderInterface, log *logrus.Logger) *Pusher {
	return &Pusher{
		ch:     make(chan *Message, cacheSize),
		sender: sender,
		log:    log,
	}
}

func (p *Pusher) Push(msg *Message) {
	p.ch <- msg
}

func (p *Pusher) Run() {
	for msg := range p.ch {
		if err := p.sender.Send(msg); err != nil {
			p.log.Errorf("Sending message error: %s", err)
		}
	}
}

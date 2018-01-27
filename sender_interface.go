package main

type SenderInterface interface {
	Send(msg *Message) error
}

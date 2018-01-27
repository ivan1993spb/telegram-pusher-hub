package main

type PusherInterface interface {
	Push(msg *Message)
}

package main

type commandId int

const (
	CMD_SUBSCRIBE commandId = iota
	CMD_SEND
)

type command struct {
	id     commandId
	client *client
	args   []string
}

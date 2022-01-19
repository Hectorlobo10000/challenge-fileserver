package main

type commandId int

const (
	CMD_RECEIVE commandId = iota
)

type command struct {
	id      commandId
	message *message
}

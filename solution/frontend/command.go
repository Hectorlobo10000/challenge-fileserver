package main

import socketio "github.com/googollee/go-socket.io"

type commandId int64

const (
	CMD_LISTENING commandId = iota
)

type command struct {
	id               commandId
	socketConnection socketio.Conn
	message          string
}

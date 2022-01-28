package main

import (
	"log"
	"net"
)

type client struct {
	commands chan command
}

func newClient() *client {
	return &client{
		commands: make(chan command),
	}
}

func initialization() net.Conn {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return conn
}

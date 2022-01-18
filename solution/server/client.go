package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	channel *channel
}

func (c *client) readInput() {
	for {
		message, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			return
		}

		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")
		command := strings.TrimSpace(args[0])

		switch command {
		case "-channel":
			log.Println("set a channel")
		}
	}
}

func (c *client) listen() {
	for {
		message, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")

		command := strings.TrimSpace(args[0])

		switch command {
		case "-channel":
			log.Println("print")
		}

	}
}

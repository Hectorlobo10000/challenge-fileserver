package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	channel  *channel
	commands chan<- command
}

func (c *client) readCommand() {
	for {
		message, err := bufio.NewReader(c.conn).ReadString('\n')
		//message, _, err := bufio.NewReader(c.conn).ReadLine()

		if err != nil {
			log.Fatalf("readCommand func: %s", err.Error())
			return
		}

		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")

		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "-subscribe":
			c.commands <- command{id: CMD_SUBSCRIBE, client: c, args: args}
		case "-send":
			c.commands <- command{id: CMD_SEND, client: c, args: args}
		default:
			log.Printf("unkown command: %s", cmd)
		}

	}
}

func (c *client) sendMessage(message string) {
	c.conn.Write([]byte(message + "\n"))
}

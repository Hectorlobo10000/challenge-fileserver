package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type message struct {
	conn     net.Conn
	commands chan<- command
}

func (m *message) readCommand() {
	for {
		message, err := bufio.NewReader(m.conn).ReadString('\n')

		if err != nil {
			return
		}

		message = strings.Trim(message, "\r\n")

		log.Println(message)

		m.commands <- command{id: CMD_RECEIVE, message: m}
	}

}

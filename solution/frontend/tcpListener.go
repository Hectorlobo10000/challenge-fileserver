package main

import (
	"bufio"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net"
	"strings"
)

type tcpServer struct {
	commands chan command
}

func newTcpServer() *tcpServer {
	return &tcpServer{
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

func (s *tcpServer) listen() {
	for command := range s.commands {
		switch command.id {
		case CMD_LISTENING:
			command.socketConnection.Emit("message", command.message)
		}
	}
}

func (s *tcpServer) tcpHandler(si socketio.Conn, conn net.Conn) {

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			return
		}

		message = strings.Trim(message, "\r\n")

		s.commands <- command{
			id: CMD_LISTENING, socketConnection: si, message: message,
		}
		log.Println(message)
	}
}

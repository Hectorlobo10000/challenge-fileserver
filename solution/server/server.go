package main

import (
	"log"
	"net"
)

type server struct {
}

func (s *server) handlerNewConnection(conn net.Conn) {
	log.Println("New client has joined: %s", conn.RemoteAddr().String())

	c := &client{conn: conn}

	c.listen()
}

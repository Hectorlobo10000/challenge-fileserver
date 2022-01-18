package main

import (
	"log"
	"net"
)

func newServer() *server {
	return &server{}
}

func initialization() net.Listener {
	listener, err := net.Listen("tcp", ":9999")

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	return listener
}

func main() {
	listener := initialization()
	server := newServer()

	defer listener.Close()
	log.Println("Server running on port: 99999")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf(err.Error())
			continue
		}

		go server.handlerNewConnection(conn)
	}
}

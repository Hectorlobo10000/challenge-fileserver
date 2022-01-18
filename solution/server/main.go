package main

import (
	"log"
)

func main() {
	listener := initialization()
	server := newServer()

	go server.listen()

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

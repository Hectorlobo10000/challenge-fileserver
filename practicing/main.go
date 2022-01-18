package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9999")

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer listener.Close()
	log.Println("Server started on port: 9999")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf(err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	bytes := make([]byte, 1000)

	message, err := conn.Read(bytes)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	log.Println("Client: ", conn.RemoteAddr().String(), "/ message: ", string(bytes[:message]))

	messageToSend := "This is a message sent from server."

	conn.Write([]byte(messageToSend))
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	go newClient()

	var input string
	fmt.Scanln(&input)
}

func newClient() {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	message := "new client connected: " + conn.RemoteAddr().String()

	conn.Write([]byte(message))

	go modeReceive(conn)

	log.Println("message sent: ", message)
}

func modeReceive(conn net.Conn) {
	bytes := make([]byte, 1000)

	message, err := conn.Read(bytes)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	log.Println("Server: ", string(bytes[:message]))
}

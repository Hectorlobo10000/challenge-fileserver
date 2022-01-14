package main

import "fmt"
import "net"
import "encoding/gob"

func handleClient(client net.Conn) {
	var message string
	err := gob.NewDecoder(client).Decode(&message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensaje: ", message)
}

func server() {
	server, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(client)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
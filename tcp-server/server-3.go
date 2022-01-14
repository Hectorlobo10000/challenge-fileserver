package main

import "fmt"
import "net"
import "encoding/gob"

type Person struct {
	Name string
	Email []string
}

func handleClient(client net.Conn) {
	var person Person
	err := gob.NewDecoder(client).Decode(&person)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensaje: ", person)
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
package main

import "fmt"
import "net"

func handleClient(client net.Conn) {
	b := make([]byte, 100)

	bytes, err := client.Read(b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensaje: ", string(b[:bytes]))
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
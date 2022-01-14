package main

import "fmt"
import "net"
import "encoding/gob"

func client() {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	message:= "Hello, world!"
	fmt.Println(message)
	err = gob.NewEncoder(conn).Encode(message)

	if err != nil {
		fmt.Println(err)		
	}

	conn.Close()
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
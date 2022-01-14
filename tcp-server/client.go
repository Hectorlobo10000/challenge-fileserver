package main

import "fmt"
import "net"

func client() {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	message:= "Hello, world!"
	fmt.Println(message)

	conn.Write([]byte(message))
	conn.Close()

}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
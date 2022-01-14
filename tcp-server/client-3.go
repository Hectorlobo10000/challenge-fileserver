package main

import "fmt"
import "net"
import "encoding/gob"

type Person struct {
	Name string
	Email []string
}

func client(person Person) {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(conn).Encode(person)

	if err != nil {
		fmt.Println(err)		
	}

	conn.Close()
}

func main() {
	person := Person {
		Name: "Hector Lobo",
		Email:[]string{
			"hector.lobo10000@gmail.com",
			"hlobo@televicentro.com",
		},
	}
	go client(person)

	var input string
	fmt.Scanln(&input)
}
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	conn := initialization()

	defer conn.Close()
	log.Println("Connected to server...")

	scanner := bufio.NewScanner(os.Stdin)

loop:
	for scanner.Scan() {

		message := scanner.Text() + "\n"
		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "-subscribe":
			print("subscribed... \n")
			conn.Write([]byte(message + "\n"))
		case "-listen":
			print("listening... \n")
			break loop
		case "-send":
			print("sending... \n")
			conn.Write([]byte(message + "\n"))
		}
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Println(err.Error())
		}

		message = strings.Trim(message, "\r\n")

		log.Println(message)
	}

}

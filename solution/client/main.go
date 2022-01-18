package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")

	defer conn.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		go handler(conn, line)
	}

}

func handler(conn net.Conn, line string) {
	conn.Write([]byte(line + "\n"))
}

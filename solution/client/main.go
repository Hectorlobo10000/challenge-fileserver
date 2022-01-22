package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"
	"path/filepath"
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
			interCommand(args, conn)
			//conn.Write([]byte(message + "\n"))
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

func interCommand(args []string, conn net.Conn) {
	mydir, _ := os.Getwd()

	filePath := filepath.Join(mydir, strings.TrimSpace(args[1]))

	file, err := os.Open(filePath)

	fileInfo, _ := file.Stat()

	buf := make([]byte, fileInfo.Size())

	bFile, _ := file.Read(buf)

	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	msg := &message{
		User:     conn.LocalAddr().String(),
		Command:  strings.TrimSpace(args[0]),
		Context:  strings.TrimSpace(args[1]),
		FileName: fileInfo.Name(),
		File:     buf[:bFile],
	}

	b, err := json.Marshal(msg)

	if err != nil {
		log.Println(err)
		return
	}

	jsonStruct := string(b)
	body := strings.TrimSpace(args[0]) + " " + jsonStruct + "\n"

	conn.Write([]byte(body))
}

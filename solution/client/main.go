package main

import (
	"bufio"
	"encoding/json"
	"github.com/gofrs/uuid"
	"io/ioutil"
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

		messageR := scanner.Text() + "\n"
		messageR = strings.Trim(messageR, "\r\n")

		args := strings.Split(messageR, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "-subscribe":
			print("subscribed... \n")
			conn.Write([]byte(messageR + "\n"))
		case "-listen":
			print("listening... \n")
			break loop
		case "-send":
			print("sending... \n")
			interCommand(args, conn)
		}
	}

	for {
		messageR, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Println(err.Error())
		}

		messageR = strings.Trim(messageR, "\r\n")

		var msg message
		json.Unmarshal([]byte(messageR), &msg)

		log.Printf("filename: %s --- filesize: %v ", msg.Filename, msg.FileSize)

		filename := msg.Id.String() + " " + msg.Filename
		errWriteFile := ioutil.WriteFile(filename, msg.File, 0644)

		if errWriteFile != nil {
			log.Println(errWriteFile)
		}

		log.Println("File saved")
	}

}

func interCommand(args []string, conn net.Conn) {
	fileExten, fileInfo, buf, bFile, err := readFile(args)

	newId, _ := uuid.NewV4()

	msg := &message{
		Id:          newId,
		User:        conn.LocalAddr().String(),
		Command:     strings.TrimSpace(args[0]),
		Context:     strings.TrimSpace(args[1]),
		Filename:    fileInfo.Name(),
		Extension:   fileExten,
		ContentType: getMimes(fileExten),
		FileSize:    fileInfo.Size(),
		File:        buf[:bFile],
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

func readFile(args []string) (string, os.FileInfo, []byte, int, error) {
	mydir, _ := os.Getwd()

	filePath := filepath.Join(mydir, strings.TrimSpace(args[1]))

	fileExten := filepath.Ext(filePath)

	file, err := os.Open(filePath)

	fileInfo, _ := file.Stat()

	buf := make([]byte, fileInfo.Size())

	bFile, _ := file.Read(buf)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return "", nil, nil, 0, err
	}
	return fileExten, fileInfo, buf, bFile, nil
}

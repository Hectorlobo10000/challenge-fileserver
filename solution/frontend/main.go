package main

import (
	"bufio"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"strings"
)

func main() {
	server := socketio.NewServer(nil)

	conn := initialization()
	tcpListener := newTcpServer()

	defer conn.Close()
	log.Println("Connected to server...")

	conn.Write([]byte("-subscribe #tree \n"))

	go tcpListener.listen()

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Printf("new client connected: %s", s.ID())

		for {
			message, err := bufio.NewReader(conn).ReadString('\n')

			if err != nil {
				log.Println(err)
			}

			message = strings.Trim(message, "\r\n")

			tcpListener.commands <- command{
				id: CMD_LISTENING, socketConnection: s, message: message,
			}
			log.Println(message)
		}

		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, message string) {
		log.Println("message from web client: ", message)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server running on port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

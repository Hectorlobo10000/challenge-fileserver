package main

import (
	"bufio"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	server := socketio.NewServer(nil)

	conn := initialization()

	defer conn.Close()
	log.Println("Connected to server...")

	conn.Write([]byte("-subscribe #tree \n"))

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")

		tcpHandler(s, conn)
		log.Printf("new client connected: %s", s.ID())
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

func initialization() net.Conn {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return conn
}

func tcpHandler(s socketio.Conn, conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			return
		}

		message = strings.Trim(message, "\r\n")

		s.Emit("message", "hoooooola")
		log.Println(message)
	}
}

package main

import (
	"log"
	"net"
)

type client struct {
	commands chan command
}

func newClient() *client {
	return &client{
		commands: make(chan command),
	}
}

func initialization() net.Conn {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return conn
}

//func (c *client) handlerConnection(conn net.Conn, line string) {
//	//m := &message{conn: conn, commands: c.commands}
//
//	conn.Write([]byte(line + "\n"))
//
//	c.listen(conn)
//
//	//m.readCommand()
//}
//
//func (c *client) listen(conn net.Conn) {
//	for {
//		message, err := bufio.NewReader(conn).ReadString('\n')
//
//		if err != nil {
//			return
//		}
//
//		message = strings.Trim(message, "\r\n")
//
//		log.Println(message)
//	}
//}

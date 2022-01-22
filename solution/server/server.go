package main

import (
	"log"
	"net"
	"strings"
)

type server struct {
	channels map[string]*channel
	commands chan command
}

func newServer() *server {
	return &server{
		channels: make(map[string]*channel),
		commands: make(chan command),
	}
}

func initialization() net.Listener {
	listener, err := net.Listen("tcp", ":9999")

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	return listener
}

func (s *server) listen() {
	for command := range s.commands {
		switch command.id {
		case CMD_SUBSCRIBE:
			s.subscribe(command.client, command.args)
		case CMD_SEND:
			s.send(command.client, command.args)
		}
	}
}

func (s *server) handlerNewConnection(conn net.Conn) {
	log.Printf("New client has joined: %s", conn.RemoteAddr().String())

	c := &client{conn: conn, name: "anonymous", commands: s.commands}

	c.readCommand()
}

func (s *server) subscribe(c *client, args []string) {
	if len(args) < 2 {
		log.Println("channel name is required.")
		return
	}

	channelName := args[1]

	ch, ok := s.channels[channelName]

	if !ok {
		ch = &channel{name: channelName, members: make(map[net.Addr]*client)}
		s.channels[channelName] = ch
	}

	ch.members[c.conn.RemoteAddr()] = c

	s.changeCurrentChannel(c)
	c.channel = ch

	log.Printf("%s joined the %s channel", c.name, c.channel.name)
}

func (s *server) changeCurrentChannel(c *client) {
	if c.channel != nil {
		delete(s.channels[c.channel.name].members, c.conn.RemoteAddr())
	}
}

func (s *server) send(c *client, args []string) {
	if len(args) < 2 {
		log.Println("message name is required.")
		return
	}

	message := strings.Join(args[1:], " ")
	//log.Printf("message sent from: %s", c.name)
	//c.channel.broadcast(c, c.name+": "+message)
	c.channel.broadcast(c, message)
	log.Printf("message sent client: %s, channel: %s", c.name, c.channel.name)
}

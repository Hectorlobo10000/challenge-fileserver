package main

import "net"

type channel struct {
	name    string
	members map[net.Addr]*client
}

func (ch *channel) broadcast(sender *client, message string) {
	for addr, mem := range ch.members {
		if sender.conn.RemoteAddr() != addr {
			mem.sendMessage(message)
		}
	}
}

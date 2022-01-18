package main

import "net"

type channel struct {
	name    string
	clients map[net.Addr]*client
}

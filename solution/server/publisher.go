package main

import "bufio"

type publisher struct {
	clients map[string]*client
}

func (p *publisher) listen() {
	for {
		message, err := bufio.NewReader()
	}
}

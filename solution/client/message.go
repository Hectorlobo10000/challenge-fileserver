package main

import "github.com/gofrs/uuid"

type message struct {
	Id          uuid.UUID `json:"id"`
	User        string    `json:"user"`
	Command     string    `json:"command"`
	Context     string    `json:"context"`
	Filename    string    `json:"filename"`
	Extension   string    `json:"extension"`
	ContentType string    `json:"contentType"`
	File        []byte    `json:"file"`
}

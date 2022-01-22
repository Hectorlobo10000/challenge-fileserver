package main

type message struct {
	User     string `json:"user"`
	Command  string `json:"command"`
	Context  string `json:"context"`
	FileName string `json:"fileName"`
	File     []byte `json:"file"`
}

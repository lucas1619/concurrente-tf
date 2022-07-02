package main

import (
	"encoding/json"
	"net"
	"os"
)

type Frame struct {
	Cmd    string   `json:"cmd"`
	Sender string   `json:"sender"`
	Data   []string `json:"data"`
}

func send(remote string, frame Frame) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		enc.Encode(frame)
	}
}

func main() {
	remote := os.Args[1]
	send(remote, Frame{"cliRegister", "test", []string{"1300", "COORPORACION GAAAAA"}})
}

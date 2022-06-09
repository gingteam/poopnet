package main

import (
	"math/rand"
	"net"
	"time"
)

var (
	botList *BotList = NewBotList()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	serverListener, err := net.Listen("tcp", SERVER_ADDRESS)
	if err != nil {
		if DEBUG_MODE {
			panic(err)
		}
		return
	}
	go BotServer()
	for {
		connection, err := serverListener.Accept()
		if err != nil {
			if DEBUG_MODE {
				panic(err)
			}
			return
		}
		server := NewServer(connection)
		go server.Handle()
	}
}

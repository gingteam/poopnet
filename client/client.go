package main

import (
	"net"
	"strings"
)

type Client struct {
	Connection net.Conn
}

func NewClient(connection net.Conn) *Client {
	return &Client{
		Connection: connection,
	}
}

func (client *Client) Send(data string) {
	dataBytes := []byte(data)
	client.Connection.Write(dataBytes)
}

func (client *Client) Recv(size int) (string, error) {
	buffer := make([]byte, size)
	if _, err := client.Connection.Read(buffer); err != nil {
		return "", err
	}
	var data string = string(buffer)
	data = strings.Replace(data, "\x00", "", -1)
	data = strings.TrimSpace(data)
	return data, nil
}

func (client *Client) Handle() {
	defer client.Connection.Close()
	client.Send(VERIFY_MESSAGE)
	for {
		data, err := client.Recv(1024)
		if err != nil {
			break
		}
		if data == PING_MESSAGE {
			client.Send(PONG_MESSAGE)
			continue
		}
		command, err := NewCommand(data)
		if err != nil {
			continue 
		}
		command.Handle()
	}
}

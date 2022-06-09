package main

const (
	CONNECT_RETRIES int    = 3
	SERVER_IP       string = "127.0.0.1"
	SERVER_PORT     string = "54321"
	SERVER_ADDRESS  string = SERVER_IP + ":" + SERVER_PORT
	VERIFY_MESSAGE  string = "0x00"
	PING_MESSAGE    string = "0x01"
	PONG_MESSAGE    string = "0x02"
	COMMAND_PREFIX  string = "!"
)

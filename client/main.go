package main

import "net"

func main() {
	retries := 0
	for {
		if retries >= CONNECT_RETRIES {
			break
		}
		conn, err := net.Dial("tcp", SERVER_ADDRESS)
		if err != nil {
			retries++
			continue
		}
		client := NewClient(conn)
		client.Handle()
	}
}

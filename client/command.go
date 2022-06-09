package main

import (
	"fmt"
	"strings"
)

type Command struct {
	Action     string
	ParsedArgs []string
	RawArgs    string
	Contents   string
}

func NewCommand(command string) (*Command, error) {
	if string(command[0]) != COMMAND_PREFIX {
		return nil, fmt.Errorf("invalid command prefix")
	}
	var contents string = command[1:]
	var parsedCommand []string = strings.Split(contents, " ")
	var action string = parsedCommand[0]
	var parsedArgs []string = parsedCommand[1:]
	var rawArgs string = strings.Join(parsedArgs, " ")
	return &Command{
		Action:     action,
		RawArgs:    rawArgs,
		ParsedArgs: parsedArgs,
		Contents:   contents,
	}, nil
}

func (command *Command) Handle() {
	var action string = command.Action
	//var parsedArgs []string = command.ParsedArgs
	var rawArgs string = command.RawArgs
	//var contents string = command.Contents
	switch action {
	case "shell":
		err := ShellExec(rawArgs)
		if err != nil {
			return
		}
	}
}

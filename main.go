package main

import (
	"log"
	"os"

	"github.com/ksahli/tsubaki/cmd"
)

func main() {
	var command cmd.Command

	switch arguments := os.Args; arguments[1] {
	case "initialize":
		command = cmd.New(arguments[1:])
	case "start":
		command = cmd.New(arguments[1:])
	default:
		log.Fatal("no command specified")
	}

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

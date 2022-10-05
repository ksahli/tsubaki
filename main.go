package main

import (
	"log"
	"os"

	"github.com/ksahli/tsubaki/cmd/initialize"
)

type Command interface {
	Execute() error
}

func main() {
	var command Command
	arguments := os.Args
	switch arguments[1] {
	case "initialize":
		command = initialize.New(arguments[1:])
	default:
		log.Fatal("no command specified")
	}
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

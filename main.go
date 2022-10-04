package main

import (
	"log"
	"os"

	"github.com/ksahli/tsubaki/cmd"
)

func main() {
	var command cmd.Command

	switch arguments := os.Args; arguments[1] {
	default:
		log.Fatal("no command was provided")
	}

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

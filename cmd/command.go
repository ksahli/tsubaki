package cmd

import (
	"flag"
)

type Command interface {
	Execute() error
}

func New(arguments []string) Command {
	switch arguments[0] {
	case "initialize":
		return initialize(arguments)
	case "start":
		return start(arguments)
	default:
		return nil
	}
}

func initialize(arguments []string) Command {
	flagSet := flag.NewFlagSet("initialize", flag.ExitOnError)

	var storage string
	flagSet.StringVar(&storage, "bbolt.storage", "", "storage location")

	flagSet.Parse(arguments[1:])
	command := Initialize{
		storage: storage,
	}
	return command
}

func start(arguments []string) Command {
	flagSet := flag.NewFlagSet("start", flag.ExitOnError)

	var storage string
	flagSet.StringVar(&storage, "bbolt.storage", "", "storgae location")

	var address string
	flagSet.StringVar(&address, "endpoints.address", "", "endpoint address")

	flagSet.Parse(arguments[1:])
	command := Start{
		storage: storage,
		address: address,
	}
	return command
}

package initialize

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ksahli/tsubaki/pkg/storage"
)

var logger = log.New(os.Stdout, " [initialize] ", log.Ldate)

type Command struct {
	storage string
	buckets [][]byte
}

func (command Command) Execute() error {
	logger.Println("initializing buckets")
	db, openErr := storage.Open(command.storage)
	if openErr != nil {
		err := fmt.Errorf("failed to initialize storage: %w", openErr)
		return err
	}
	storage.Setup(db, storage.Buckets)
	logger.Println("all buckets have been initialized")
	return nil
}

func New(arguments []string) *Command {
	command := new(Command)
	flagSet := flag.NewFlagSet("initialize", flag.ExitOnError)
	flagSet.StringVar(&command.storage, "bbolt.storage", "", "storage location")
	flagSet.Parse(arguments[1:])
	return command
}

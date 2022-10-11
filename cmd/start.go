package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ksahli/tsubaki/pkg/endpoints"
	"github.com/ksahli/tsubaki/pkg/storage"
)

type Start struct {
	storage string
	address string
}

func (command Start) Execute() error {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	db, openErr := storage.Open(command.storage)
	if openErr != nil {
		err := fmt.Errorf("failed to execute command: %w", openErr)
		return err
	}
	defer db.Close()

	router := endpoints.Router()

	server := endpoints.New(command.address, router)
	go server.Start(nil)

	<-exit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if stopErr := server.Stop(ctx); stopErr != nil {
		err := fmt.Errorf("failed to stop command: %w", stopErr)
		return err
	}
	return nil
}

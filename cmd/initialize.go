package cmd

import (
	"fmt"

	"github.com/ksahli/tsubaki/pkg/storage"
)

type Initialize struct {
	storage string
}

func (command Initialize) Execute() error {
	db, openErr := storage.Open(command.storage)
	if openErr != nil {
		err := fmt.Errorf("failed to initialize storage: %w", openErr)
		return err
	}
	storage.Setup(db, storage.Buckets)
	return nil
}

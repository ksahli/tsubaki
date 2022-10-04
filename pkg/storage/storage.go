package storage

import (
	"fmt"
	"log"
	"os"

	"go.etcd.io/bbolt"
)

var logger = log.New(os.Stdout, " [storage] ", log.Ldate)

type Database interface {
	Close() error
}

func Open(path string) (Database, error) {
	logger.Printf("opening database on '%s' ...", path)
	db, openErr := bbolt.Open(path, 0644, nil)
	if openErr != nil {
		err := fmt.Errorf("failed to open database: %w", openErr)
		return nil, err
	}
	log.Println("database open")
	return db, nil
}


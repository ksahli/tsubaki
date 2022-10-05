package storage

import (
	"fmt"
	"log"
	"os"

	"go.etcd.io/bbolt"
)

var logger = log.New(os.Stdout, " [storage] ", log.Ldate)

var Buckets = [][]byte{}

func Open(path string) (*bbolt.DB, error) {
	logger.Printf("opening database on '%s' ...", path)
	db, openErr := bbolt.Open(path, 0644, nil)
	if openErr != nil {
		err := fmt.Errorf("failed to open database: %w", openErr)
		return nil, err
	}
	logger.Println("database open")
	return db, nil
}

func Setup(db *bbolt.DB, buckets [][]byte) error {
	setup := func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
				return err
			}
		}
		return nil
	}
	return db.Update(setup)
}

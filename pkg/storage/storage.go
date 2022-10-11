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
	defer logger.Println("database open")

	db, openErr := bbolt.Open(path, 0644, nil)
	if openErr != nil {
		err := fmt.Errorf("failed to open database: %w", openErr)
		return nil, err
	}
	return db, nil
}

func Setup(db *bbolt.DB, buckets [][]byte) error {
	logger.Println("creating buckets ...")
	defer logger.Println("all buckets have been created")

	setup := func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			logger.Printf("creating bucket %s", bucket)
			if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
				return err
			}
			logger.Printf("bucket %s have been created", bucket)
		}
		return nil
	}

	return db.Update(setup)
}

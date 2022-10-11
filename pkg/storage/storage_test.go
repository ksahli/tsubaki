package storage_test

import (
	"path/filepath"
	"testing"

	"github.com/ksahli/tsubaki/pkg/storage"
	"go.etcd.io/bbolt"
)

func TestOpen(t *testing.T) {

	t.Run("failure: invalid path", func(t *testing.T) {
		db, openErr := storage.Open("")
		if openErr == nil {
			t.Fatalf("expecting an error for, got nothing")
		}

		if db != nil {
			db.Close()
			t.Fatalf("expecting databse to be nil, got %#v", db)
		}
	})

	t.Run("success", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "test.db")
		db, openErr := storage.Open(path)
		if openErr != nil {
			t.Fatalf("unexpected error %#v", openErr)
		}

		defer db.Close()
		if db == nil {
			t.Fatal("expecting a database, got nothing")
		}
	})

}

func TestSetup(t *testing.T) {

	t.Run("failure: nil buckets", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "test.db")
		db, openErr := storage.Open(path)
		if openErr != nil {
			t.Fatalf("unexpected error %#v", openErr)
		}

		defer db.Close()
		buckets := [][]byte{nil}
		if err := storage.Setup(db, buckets); err == nil {
			t.Fatal("expecting an error, got nothing")
		}
	})

	t.Run("success", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "test.db")
		db, openErr := storage.Open(path)
		if openErr != nil {
			t.Fatalf("unexpected error %#v", openErr)
		}

		defer db.Close()
		buckets := [][]byte{
			[]byte("bucket_1"), []byte("bucket_2"), []byte("bucket_3"),
		}
		if err := storage.Setup(db, buckets); err != nil {
			t.Fatalf("unexpected error %#v", err)
		}

		for _, bucket := range buckets {
			check := func(tx *bbolt.Tx) error {
				if tx.Bucket(bucket) == nil {
					t.Fatalf("expecting '%s' to have been created", bucket)
				}
				return nil
			}
			if err := db.View(check); err != nil {
				t.Fatalf("unexpected error %#v", err)
			}
		}
	})

}

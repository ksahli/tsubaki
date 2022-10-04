package storage_test

import (
	"path/filepath"
	"testing"

	"github.com/ksahli/tsubaki/pkg/storage"
)

func TestOpen(t *testing.T) {

	t.Run("failure: invalid path", func(t *testing.T) {
		db, err := storage.Open("")
		if err == nil {
			t.Fatalf("expecting an error for, got nothing")
		}

		if db != nil {
			db.Close()
			t.Fatalf("expecting databse to be nil, got %#v", db)
		}
	})

	t.Run("success", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "test.db")
		db, err := storage.Open(path) 
		if err != nil {
			t.Fatalf("unexpected error %#v", err)
		}

		defer db.Close()
		if db == nil {
			t.Fatal("expecting a database, got nothing")
		}
	})

}


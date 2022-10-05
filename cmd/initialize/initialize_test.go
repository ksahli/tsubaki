package initialize_test

import (
	"testing"

	"github.com/ksahli/tsubaki/cmd/initialize"
)

func TestCommandExecute(t *testing.T) {

	t.Run("failure: invalid storage path", func(t *testing.T) {
		arguments := []string{"initialize"}
		command := initialize.New(arguments)
		if err := command.Execute(); err == nil {
			t.Fatal("expecting an error, got nothing")
		}
	})

	t.Run("success", func(t *testing.T) {
		arguments := []string{"initialize", "-bbolt.storage=tsubaki.db"}
		command := initialize.New(arguments)
		if err := command.Execute(); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

}

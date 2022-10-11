package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/ksahli/tsubaki/pkg/endpoints"
)

func TestNew(t *testing.T) {
	middleware := func(http.Handler) http.Handler {
		return nil
	}

	router := endpoints.Router(middleware)
	if router == nil {
		t.Fatal("expecting a router, got nothing")
	}
}

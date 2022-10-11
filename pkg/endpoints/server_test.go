package endpoints_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/ksahli/tsubaki/pkg/endpoints"
)

func TestStart(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		payload := []byte("test payload")
		handlerFunc := func(w http.ResponseWriter, _ *http.Request) {
			w.Write(payload)
		}
		handler := http.HandlerFunc(handlerFunc)

		server := endpoints.New("localhost:8080", handler)

		errchan := make(chan error, 1)
		go server.Start(errchan)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		defer server.Stop(ctx)

		request, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		client := http.Client{}
		response, err := client.Do(request)
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
		defer response.Body.Close()

		got, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if expected := payload; !bytes.Equal(expected, got) {
			t.Fatalf("expecting %s, got %s", expected, got)
		}

		close(errchan)
		for err := range errchan {
			if err != nil {
				t.Fatalf("unexpected error: %#v", err)
			}
		}
	})

	t.Run("failure: invalid address", func(t *testing.T) {
		server := endpoints.New("invalid address", nil)

		errchan := make(chan error, 1)
		server.Start(errchan)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		defer server.Stop(ctx)

		errs := make([]error, 0)
		close(errchan)
		for err := range errchan {
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) == 0 {
			t.Fatal("expecting errors, got nothing")
		}
	})

}

func TestStop(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		server := endpoints.New("localhost:8080", nil)
		go server.Start(nil)

		deadline := time.Now().Add(10 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("failure", func(t *testing.T) {
		handlerFunc := func(http.ResponseWriter, *http.Request) {
			time.Sleep(1 * time.Hour)
		}
		handler := http.HandlerFunc(handlerFunc)

		server := endpoints.New("localhost:8080", handler)
		go server.Start(nil)

		request, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
		client := http.Client{}
		go client.Do(request)

		time.Sleep(1 * time.Second)

		deadline := time.Now().Add(1 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		if err := server.Stop(ctx); err == nil {
			t.Fatalf("expecting an error, got nothing: %#v", ctx.Err())
		}
	})

}

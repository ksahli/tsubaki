package endpoints

import (
	"context"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "[endpoints] ", log.Ldate)

type Server struct {
	server *http.Server
}

func (endpoints *Server) Start(errchan chan error) {
	logger.Println("starting http server ...")
	defer logger.Println("http server started")

	err := endpoints.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		errchan <- err
	}
}

func (endpoints *Server) Stop(ctx context.Context) error {
	logger.Println("stopping server ...")
	defer logger.Println("server stopped")

	err := endpoints.server.Shutdown(ctx)
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func New(address string, handler http.Handler) *Server {
	server := &http.Server{
		Handler: handler,
		Addr:    address,
	}
	endpoints := Server{server: server}
	return &endpoints
}

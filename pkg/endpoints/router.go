package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(middlewares ...mux.MiddlewareFunc) http.Handler {
	router := mux.NewRouter().
		PathPrefix("/api/v1/").
		Headers("Content-Type", "application/json").
		Subrouter()

	for _, middleware := range middlewares {
		middlewareFunc := mux.MiddlewareFunc(middleware)
		router.Use(middlewareFunc)
	}

	return router
}


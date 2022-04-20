package main

import (
	"exampleModule/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler{
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}

package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit HomePage Get Endpoint.")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit AboutPage Get Endpoint.")
}

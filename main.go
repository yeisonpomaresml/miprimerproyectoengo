package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"www.github.com/miprimerproyectoengo/handlers"
)

func main() {
	usuarioController := handlers.NewUsuarioController()

	router := mux.NewRouter()

	router.HandleFunc("/", usuarioController.HandleGeneral)
	router.HandleFunc("/{id}", usuarioController.HandleOne)

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	server.ListenAndServe()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhuiyanmarzan/eghissab/internal/config"
	"github.com/bhuiyanmarzan/eghissab/internal/handlers"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Printf("Starting server on port %s\n", cfg.Port)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.Health)
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

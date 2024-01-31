package main

import (
	"h4g-vms/backend/pkg/config"
	"h4g-vms/backend/pkg/handlers"
	"h4g-vms/backend/pkg/store"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Initialize the store
    dbStore, err := store.New(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to create store: %v", err)
    }

    // Create a router
    r := chi.NewRouter()

    // Setup routes
    handlers.SetupRoutes(r, dbStore)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}

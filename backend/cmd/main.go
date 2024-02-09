package main

import (
	"h4g-vms/pkg/handlers"
	"h4g-vms/pkg/store"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
    dbStore, _, err := store.Open()
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }

    r := chi.NewRouter()

    handlers.SetupRoutes(r, dbStore)

    log.Printf("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))


}

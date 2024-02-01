package handlers

import (
	"h4g-vms/pkg/store"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(r *chi.Mux, dbStore *store.Store) {
    r.Get("/items", getItemsHandler(dbStore))
    // Add more routes here
}

func getItemsHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        items, err := dbStore.GetItems()
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
		// Write items to response
    }
}

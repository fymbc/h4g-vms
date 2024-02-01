package handlers

import (
	"encoding/json"
	"h4g-vms/pkg/models"
	"h4g-vms/pkg/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func SetupRoutes(r *chi.Mux, dbStore *store.Store) {
    r.Get("/items", getAllItemsHandler(dbStore))
    r.Get("/items/{id}", getItemHandler(dbStore))
    r.Post("/items", createItemHandler(dbStore))
    r.Put("/items/{id}", updateItemHandler(dbStore))
    r.Delete("/items/{id}", deleteItemHandler(dbStore))
}

func getAllItemsHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        items, err := dbStore.GetAllItems()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(items)
    }
}

func getItemHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid item ID", http.StatusBadRequest)
            return
        }
        item, err := dbStore.GetItem(id)
        if err != nil {
            http.Error(w, "Item not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(item)
    }
}

func createItemHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var item models.Item
        err := json.NewDecoder(r.Body).Decode(&item)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        err = dbStore.CreateItem(&item)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(item)
    }
}

func updateItemHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid item ID", http.StatusBadRequest)
            return
        }
        var item models.Item
        err = json.NewDecoder(r.Body).Decode(&item)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        item.ID = id
        err = dbStore.UpdateItem(&item)
        if err != nil {
            http.Error(w, "Item not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(item)
    }
}

func deleteItemHandler(dbStore *store.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid item ID", http.StatusBadRequest)
            return
        }
        err = dbStore.DeleteItem(id)
        if err != nil {
            http.Error(w, "Item not found", http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}

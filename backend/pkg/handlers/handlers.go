package handlers

import (
	"h4g-vms/pkg/handlers/activities"
	authentication "h4g-vms/pkg/handlers/authentication"
	"h4g-vms/pkg/store"

	"github.com/go-chi/chi"
)

func SetupRoutes(r *chi.Mux, dbStore *store.Store) {
    // Public routes
	r.Route("/", func(r chi.Router) {
		r.Get("/login", authentication.LogIn(dbStore))
		r.Get("/signup", authentication.SignUp(dbStore))
		// r.Route("/enrol", func(r chi.Router) {
		// 	r.Get("/{id}", enrolHandler(dbStore))
		// })
		// r.Post("/feedback", feedbackHandler(dbStore))
	})

    r.Route("/activities", func(r chi.Router) {
        r.Get("/", activities.HandleList(dbStore))
        r.Post("/", activities.HandleCreate(dbStore))
        r.Route("/{id}", func(r chi.Router) {
            r.Get("/{id}", activities.HandleRead(dbStore))
            r.Put("/", activities.HandleUpdate(dbStore))
            r.Delete("/", activities.HandleDelete(dbStore))
		})

    })

    // Admin routes
	// r.Route("/admin", func(r chi.Router) {
	// 	r.Use(AdminOnly) // Middleware to check if the user is an admin
	// 	r.Get("/dashboard", dashboardHandler(dbStore))
	// 	r.Get("/create", createHandler(dbStore))
	// 	r.Get("/cert", certHandler(dbStore))
	// 	r.Get("/volunteers", volunteersHandler(dbStore))
	// })
}


// // AdminOnly is a middleware function to check if the user is an admin.
// func AdminOnly(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Check if the user is an admin
// 		// If not, return an unauthorized error, for example:
// 		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		// Otherwise, call next.ServeHTTP(w, r)
// 	})
// }

// // Implement the rest of the handler functions as needed...

// func getAllItemsHandler(dbStore *store.Store) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         items, err := dbStore.GetAllItems()
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }
//         json.NewEncoder(w).Encode(items)
//     }
// }

// func getItemHandler(dbStore *store.Store) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         id, err := strconv.Atoi(chi.URLParam(r, "id"))
//         if err != nil {
//             http.Error(w, "Invalid item ID", http.StatusBadRequest)
//             return
//         }
//         item, err := dbStore.GetItem(id)
//         if err != nil {
//             http.Error(w, "Item not found", http.StatusNotFound)
//             return
//         }
//         json.NewEncoder(w).Encode(item)
//     }
// }

// func createItemHandler(dbStore *store.Store) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         var item models.Item
//         err := json.NewDecoder(r.Body).Decode(&item)
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusBadRequest)
//             return
//         }
//         err = dbStore.CreateItem(&item)
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }
//         w.WriteHeader(http.StatusCreated)
//         json.NewEncoder(w).Encode(item)
//     }
// }

// func updateItemHandler(dbStore *store.Store) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         id, err := strconv.Atoi(chi.URLParam(r, "id"))
//         if err != nil {
//             http.Error(w, "Invalid item ID", http.StatusBadRequest)
//             return
//         }
//         var item models.Item
//         err = json.NewDecoder(r.Body).Decode(&item)
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusBadRequest)
//             return
//         }
//         item.ID = id
//         err = dbStore.UpdateItem(&item)
//         if err != nil {
//             http.Error(w, "Item not found", http.StatusNotFound)
//             return
//         }
//         json.NewEncoder(w).Encode(item)
//     }
// }

// func deleteItemHandler(dbStore *store.Store) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         id, err := strconv.Atoi(chi.URLParam(r, "id"))
//         if err != nil {
//             http.Error(w, "Invalid item ID", http.StatusBadRequest)
//             return
//         }
//         err = dbStore.DeleteItem(id)
//         if err != nil {
//             http.Error(w, "Item not found", http.StatusNotFound)
//             return
//         }
//         w.WriteHeader(http.StatusOK)
//     }
// }

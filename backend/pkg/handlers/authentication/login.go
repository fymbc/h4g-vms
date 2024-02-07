package authentication

import (
	"errors"
	"h4g-vms/pkg/dataaccess/users"
	"h4g-vms/pkg/json"
	"h4g-vms/pkg/store"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LogInParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}



func LogIn(dbStore *store.Store ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		var logInParams LogInParams
		err := json.DecodeParams(r.Body, &logInParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := users.Read(db, logInParams.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		userNotFound := user.ID == 0
		if userNotFound {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(logInParams.Password))

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Failed to compare password", http.StatusInternalServerError)
			return
		}

		response := LoginResponse{
			Success: true,
			Message: "Login successful",
		}

		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

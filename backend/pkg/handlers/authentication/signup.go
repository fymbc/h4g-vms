package authentication

import (
	"fmt"
	"h4g-vms/pkg/dataaccess/users"
	"h4g-vms/pkg/json"
	"h4g-vms/pkg/models"
	"h4g-vms/pkg/store"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignUpParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func SignUp(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		var signUpParams SignUpParams
		err := json.DecodeParams(r.Body, &signUpParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpParams.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
			return
		}

		user := models.User{
			Email: signUpParams.Email,
			EncryptedPassword: string(encryptedPassword),
		}

		existingUser, _ := users.Read(db, user.Email)
		if existingUser.ID != 0 {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		createdUser, err := users.Create(db, &user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		response := SignUpResponse{
			Success: true,
			Message: fmt.Sprintf("User \"%s\" created", createdUser.Email),
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

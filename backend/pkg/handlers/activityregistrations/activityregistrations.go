package activityregistrations

import (
	"h4g-vms/pkg/dataaccess/activityregistrations"
	"h4g-vms/pkg/json"
	"h4g-vms/pkg/models"
	"h4g-vms/pkg/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ActivityRegistrationParams struct {
	UserID     int `json:"user_id"`
	ActivityID int `json:"activity_id"`
}

func (p *ActivityRegistrationParams) ToModel() *models.ActivityRegistration {
	return &models.ActivityRegistration{
		UserID:     p.UserID,
		ActivityID: p.ActivityID,
	}
}

func HandleCreate(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		var params ActivityRegistrationParams
		err := json.DecodeParams(r.Body, &params)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
		activityRegistration := params.ToModel()
		err = activityregistrations.Create(db, activityRegistration)
		if err != nil {
			http.Error(w, "Failed to create activity registration", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(activityRegistration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func HandleDelete(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		activityRegistrationID, err := strconv.Atoi(chi.URLParam(r, "activityRegistrationID"))
		if err != nil {
			http.Error(w, "Failed to parse activity registration ID", http.StatusBadRequest)
			return
		}
		err = activityregistrations.Delete(db, activityRegistrationID)
		if err != nil {
			http.Error(w, "Failed to delete activity registration", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

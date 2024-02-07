package activities

import (
	"h4g-vms/pkg/dataaccess/activities"
	"h4g-vms/pkg/json"
	"h4g-vms/pkg/models"
	"h4g-vms/pkg/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CreateParams struct {
	baseParams
}

type UpdateParams struct {
	ID int `json:"id"`
	baseParams
}

type baseParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p *CreateParams) ToModel() *models.Activity {
	return p.baseParams.ToModel()
}

func (p *UpdateParams) ToModel() *models.Activity {
	model := p.baseParams.ToModel()
	model.ID = p.ID
	return model
}

func (p *baseParams) ToModel() *models.Activity {
	return &models.Activity{
		Name:        p.Name,
		Description: p.Description,
	}
}

func HandleList(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		activitiesList, err := activities.List(db)
		if err != nil {
			http.Error(w, "Failed to list activities", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(activitiesList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func HandleRead(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		param := chi.URLParam(r, "activityId")
		activityId, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, "Invalid activity ID", http.StatusBadRequest)
			return
		}
		activity, err := activities.Read(db, activityId)
		if err != nil {
			http.Error(w, "Failed to read activity", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(activity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func HandleCreate(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		var activity models.Activity
		err := json.DecodeParams(r.Body, &activity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdActivity, err := activities.Create(db, &activity)
		if err != nil {
			http.Error(w, "Failed to create activity", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(createdActivity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func HandleUpdate(dbStore *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := dbStore.DB
		param := chi.URLParam(r, "activityId")
		activityId, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, "Invalid activity ID", http.StatusBadRequest)
			return
		}

		var updateParams UpdateParams
		err = json.DecodeParams(r.Body, &updateParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if updateParams.ID != (activityId) {
			http.Error(w, "Mismatched activity IDs", http.StatusBadRequest)
			return
		}
		activity := updateParams.ToModel()
		activity, err = activities.Update(db, activity)
		if err != nil {
			http.Error(w, "Failed to update activity", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.EncodeView(activity)
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
		param := chi.URLParam(r, "activityId")
		activityId, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, "Invalid activity ID", http.StatusBadRequest)
			return
		}
		err = activities.Delete(db, activityId)
		if err != nil {
			http.Error(w, "Failed to delete activity", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

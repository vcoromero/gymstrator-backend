package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
	"github.com/vcoromero/gymstration-backend/src/services"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response models.ResponseAPI

	roles, err := services.GetRoles(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Roles retrieved successfully"
	response.Data = roles

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response models.ResponseAPI

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	role, err := services.GetRole(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Role retrieved successfully"
	response.Data = role

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	_, err := services.CreateRole(ctx, input.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Role created successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return

	}
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	_, err := services.UpdateRole(ctx, id, input.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Role updated successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	err := services.DeleteRole(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Role deleted successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

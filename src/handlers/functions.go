package handlers

import (
	"encoding/json"
	"net/http"

<<<<<<< HEAD
	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
	"github.com/vcoromero/gymstration-backend/src/services"
)

var input struct {
	Name       string `json:"name"`
	Descripion string `json:"description"`
}

func GetFunctions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response models.ResponseAPI

	functions, err := services.GetFunctions(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Function retrieved successfully"
	response.Data = functions

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func GetFunction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var response models.ResponseAPI

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "id" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	function, err := services.GetFunction(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Function retrieved successfully"
	response.Data = function

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func CreateFunction(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	_, err := services.CreateFunction(ctx, input.Name, input.Descripion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Function created successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func UpdateFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "id" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	_, err := services.UpdateFunction(ctx, id, input.Name, input.Descripion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Function updated successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func DeleteFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "id" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	var response models.ResponseAPI

	err := services.DeleteFunction(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Function deleted successfully"
=======
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func GetFunction(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "First function endpoint done!"
>>>>>>> 67b1c8c (improvements for role endpoints)

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

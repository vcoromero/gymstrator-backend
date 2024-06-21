package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func GetFunctions(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "Function retrieved successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func GetFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "id" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var response models.ResponseAPI

	response.Message = "Function retrieved successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func CreateFunction(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

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

	var response models.ResponseAPI

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

	var response models.ResponseAPI

	response.Message = "Function deleted successfully"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

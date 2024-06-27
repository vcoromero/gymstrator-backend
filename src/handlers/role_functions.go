package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func GetFunctionsByRole(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "GetFunctionsByRole"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func AddFunctionToRole(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "CreateRoleFunction"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func RemoveFunctionOnRole(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "DeleteRoleFunction"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func AddManyFunctionsToRole(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "CreateMultipleRoleFunction"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

func RemoveManyFunctionsOnRole(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "DeleteMultipleRoleFunction"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func GetFunction(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseAPI

	response.Message = "First function endpoint done!"

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(response)
}

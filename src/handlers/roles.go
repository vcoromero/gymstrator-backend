package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vcoromero/gymstration-backend/src/services"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := services.GetRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

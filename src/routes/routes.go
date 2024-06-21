package routes

import (
	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/src/handlers"
)

func RegisterRoutes(r *mux.Router) {
	apiRouterV1 := r.PathPrefix("/api/v1").Subrouter()

	apiRouterV1.HandleFunc("/roles/{id}", handlers.GetRoleById).Methods("GET")
	apiRouterV1.HandleFunc("/roles", handlers.GetRoles).Methods("GET")
	apiRouterV1.HandleFunc("/roles", handlers.CreateRole).Methods("POST")
	// apiRouterV1.HandleFunc("/roles/{id}", handlers.UpdateRole).Methods("PUT")
	// apiRouterV1.HandleFunc("/roles/{id}", handlers.DeleteRole).Methods("DELETE")
}

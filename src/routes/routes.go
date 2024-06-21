package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/src/handlers"
)

func RegisterRoutes(r *mux.Router) {
	apiRouterV1 := r.PathPrefix("/api/v1").Subrouter()

	apiRouterV1.HandleFunc("/roles/{id}", handlers.GetRoleById).Methods("GET")
	apiRouterV1.HandleFunc("/roles", handlers.GetRoles).Methods("GET")
	apiRouterV1.HandleFunc("/roles", handlers.CreateRole).Methods("POST")
	apiRouterV1.HandleFunc("/roles/{id}", handlers.UpdateRole).Methods("PUT")

	logRoutes(r)
	// apiRouterV1.HandleFunc("/roles/{id}", handlers.DeleteRole).Methods("DELETE")
}

func logRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			log.Printf("Route: %s Methods: NONE\n", t)
			return nil
		}

		log.Printf("Route: %s Methods: %v\n", t, methods)
		return nil
	})
	if err != nil {
		log.Fatalf("Error listing routes: %v", err)
	}
}

package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/src/handlers"
)

func RegisterRoutes(r *mux.Router) {
	apiRouterV1 := r.PathPrefix("/api/v1").Subrouter()
	rolesRouter := apiRouterV1.PathPrefix("/roles").Subrouter()
	functionsRouter := apiRouterV1.PathPrefix("/functions").Subrouter()
	roleFunctionsRouter := apiRouterV1.PathPrefix("/role-functions").Subrouter()

	rolesRouter.HandleFunc("/{id}", handlers.GetRole).Methods("GET")
	rolesRouter.HandleFunc("", handlers.GetRoles).Methods("GET")
	rolesRouter.HandleFunc("", handlers.CreateRole).Methods("POST")
	rolesRouter.HandleFunc("/{id}", handlers.UpdateRole).Methods("PUT")
	rolesRouter.HandleFunc("/{id}", handlers.DeleteRole).Methods("DELETE")

	functionsRouter.HandleFunc("/{id}", handlers.GetFunction).Methods("GET")
	functionsRouter.HandleFunc("", handlers.GetFunctions).Methods("GET")
	functionsRouter.HandleFunc("", handlers.CreateFunction).Methods("POST")
	functionsRouter.HandleFunc("/{id}", handlers.UpdateFunction).Methods("PUT")
	functionsRouter.HandleFunc("/{id}", handlers.DeleteFunction).Methods("DELETE")

	roleFunctionsRouter.HandleFunc("/{id}", handlers.GetFunctionsByRole).Methods("GET")
	roleFunctionsRouter.HandleFunc("/{id}/many", handlers.AddManyFunctionsToRole).Methods("POST")
	roleFunctionsRouter.HandleFunc("/{id}", handlers.AddFunctionToRole).Methods("POST")
	roleFunctionsRouter.HandleFunc("/{id}/many", handlers.RemoveManyFunctionsOnRole).Methods("DELETE")
	roleFunctionsRouter.HandleFunc("/{id}", handlers.RemoveFunctionOnRole).Methods("DELETE")

	logRoutes(r)
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

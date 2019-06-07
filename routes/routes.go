package routes

import (
	"go-ms-adapter/handlers"

	"github.com/gorilla/mux"
)

// Router - Api routes
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/generic", handlers.GenericHandler).Methods("POST")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	return router
}

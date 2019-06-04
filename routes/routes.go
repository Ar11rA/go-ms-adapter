package routes

import (
	"go/go-adapter-framework/handlers"

	"github.com/gorilla/mux"
)

// Router - Api routes
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/generic", handlers.GenericHandler).Methods("POST")
	return router
}

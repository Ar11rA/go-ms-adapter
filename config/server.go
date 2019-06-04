package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server object
type Server struct {
	Router *mux.Router
}

// Initialize the server
func (s *Server) Initialize(router *mux.Router) {
	s.Router = router
}

// Run server
func (s *Server) Run() {
	log.Println("Starting server at 8081")
	log.Fatal(http.ListenAndServe(":8081", s.Router))
}

package config

import (
	"encoding/json"
	"go-ms-adapter/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Server object
type Server struct {
	Router *mux.Router
}

// Configs ...
var Configs JSONConfig

// Templates ...
var Templates RequestTemplate

func setConfigs(files []string) {
	for _, file := range files {
		fileBytes := utils.ReadContents(file)
		fileContent := string(fileBytes)
		var jsonConfig TemplateParams
		fileName := strings.Split(file, "/")
		apiKey := strings.Split(fileName[2], ".")
		json.Unmarshal([]byte(fileContent), &jsonConfig)
		Configs[apiKey[0]] = jsonConfig
	}
}

func setTemplates(files []string) {
	for _, file := range files {
		fileBytes := utils.ReadContents(file)
		fileName := strings.Split(file, "/")
		apiKey := strings.Split(fileName[2], ".")
		Templates[apiKey[0]] = fileBytes
	}
}

// Initialize the server
func (s *Server) Initialize(router *mux.Router) {
	Configs = make(JSONConfig)
	Templates = make(RequestTemplate)
	configs, _ := utils.FilePathWalkDir("./resources/json_configs/")
	templates, _ := utils.FilePathWalkDir("./resources/request_templates/")
	setConfigs(configs)
	setTemplates(templates)
	s.Router = router
}

// Run server
func (s *Server) Run() {
	log.Println("Starting server at 8081")
	log.Fatal(http.ListenAndServe(":8081", s.Router))
}

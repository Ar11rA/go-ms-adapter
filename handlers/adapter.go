package handlers

import (
	"encoding/json"
	"go/go-adapter-framework/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// GenericHandler : Resource key based handler
func GenericHandler(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(string(bodyBuffer)), &result)
	log.Println("Resource Key", result["resourceKey"])
	path := "./resources/request_templates/sample_request.txt"
	fileContents := utils.ReadContents(path)
	requestTemplate := utils.FormRequest(fileContents, result)
	out := make(chan string)
	go utils.MakePostRequest("http://localhost:3000/ping", requestTemplate, out)
	utils.JSON(w, http.StatusOK, <-out)
}

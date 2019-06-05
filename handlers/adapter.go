package handlers

import (
	"encoding/json"
	"go-ms-adapter/config"
	"go-ms-adapter/utils"
	"io/ioutil"
	"net/http"
)

// GenericHandler : Resource key based handler
func GenericHandler(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)

	//request
	var result map[string]interface{}
	json.Unmarshal([]byte(string(bodyBuffer)), &result)
	resourceKey := result["resourceKey"].(string)

	requestTemplateName := config.Configs[resourceKey].RequestTemplate
	method := config.Configs[resourceKey].Method
	requestTemplateConfig := config.Templates[requestTemplateName]

	requestTemplate := utils.FormRequest(requestTemplateConfig, result)
	out := make(chan string)
	url := config.Configs[resourceKey].URL

	go utils.MakeRemoteRequest(url, method, requestTemplate, out)

	utils.JSON(w, http.StatusOK, <-out)
}

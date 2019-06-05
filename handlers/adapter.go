package handlers

import (
	"encoding/json"
	"go/go-adapter-framework/config"
	"go/go-adapter-framework/utils"
	"io/ioutil"
	"net/http"
)

// GenericHandler : Resource key based handler
func GenericHandler(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(string(bodyBuffer)), &result)
	resourceKey := result["resourceKey"].(string)
	url := config.Configs[resourceKey].URL
	requestTemplateName := config.Configs[resourceKey].RequestTemplate
	requestTemplateConfig := config.Templates[requestTemplateName]
	requestTemplate := utils.FormRequest(requestTemplateConfig, result)
	out := make(chan string)
	go utils.MakePostRequest(url, requestTemplate, out)
	utils.JSON(w, http.StatusOK, <-out)
}

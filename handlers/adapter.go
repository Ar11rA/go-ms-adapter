package handlers

import (
	"encoding/json"
	"go-ms-adapter/config"
	"go-ms-adapter/services"
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
	isRequestValid, errMessage := services.RequestValidator(config.Configs[resourceKey].RequestParams, result)
	if !isRequestValid {
		utils.Error(w, http.StatusBadRequest, errMessage)
		return
	}
	requestTemplateName := config.Configs[resourceKey].RequestTemplate
	method := config.Configs[resourceKey].Method
	requestTemplateConfig := config.Templates[requestTemplateName]

	requestTemplate := services.FormRequest(requestTemplateConfig, result)
	out := make(chan string)
	url := config.Configs[resourceKey].URL

	go services.MakeRemoteRequest(url, method, requestTemplate, out)

	utils.JSON(w, http.StatusOK, <-out)
}

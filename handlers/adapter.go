package handlers

import (
	"encoding/json"
	"go-ms-adapter/config"
	"go-ms-adapter/services"
	"go-ms-adapter/utils"
	"go-ms-adapter/constants"
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
	delete(result, "resourceKey")

	isRequestValid, errMessage := services.RequestValidator(config.Configs[resourceKey].RequestParams, result)
	if !isRequestValid {
		utils.Error(w, http.StatusBadRequest, errMessage)
		return
	}
	requestTemplateName := config.Configs[resourceKey].RequestTemplate
	method := config.Configs[resourceKey].Method
	requestTemplateConfig := config.Templates[requestTemplateName]

	queryParams := config.Configs[resourceKey].QueryParams
	requestTemplate := services.FormRequest(requestTemplateConfig, result)

	url := config.Configs[resourceKey].URL + services.FormQueryParams(method, result, queryParams)
	out := services.MakeRemoteRequest(url, method, requestTemplate)

	if _, ok := constants.ERROR_MESSAGES[out]; ok {
		utils.Error(w, http.StatusBadRequest, constants.ERROR_MESSAGES[out])
		return
	}
	utils.JSON(w, http.StatusOK, out)
}

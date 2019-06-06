package handlers

import (
	"encoding/json"
	"go-ms-adapter/config"
	"go-ms-adapter/services"
	"go-ms-adapter/utils"
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
)

// convert query params to append to url
func formQueryParams(method string, buff map[string]interface {}, queryParams []config.Params) string{
	urlQueryParams := ""

	if(len(queryParams) > 0){
		var parameters []string
		for _, element := range queryParams {
			if(buff[element.Name] != nil) {
				parameters = append(parameters, fmt.Sprintf("%s=%s", element.Name, buff[element.Name]))
			}
		}
		return "?" + strings.Join(parameters, "&")
	}
	return urlQueryParams
}

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
	requestTemplate := utils.FormRequest(requestTemplateConfig, result)

	out := make(chan string)
	url := config.Configs[resourceKey].URL + formQueryParams(method, result, queryParams)
	go utils.MakeRemoteRequest(url, method, requestTemplate, out)

	utils.JSON(w, http.StatusOK, <-out)
}

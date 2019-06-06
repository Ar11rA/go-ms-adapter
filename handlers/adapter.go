package handlers

import (
	"go-ms-adapter/config"
	"go-ms-adapter/constants"
	"go-ms-adapter/services"
	"go-ms-adapter/utils"
	"net/http"
	"reflect"
)

// GenericHandler : Resource key based handler
func GenericHandler(w http.ResponseWriter, r *http.Request) {
	requestPayload, resourceKey := utils.ParseRequest(r)
	if reflect.DeepEqual(config.Configs[resourceKey], config.TemplateParams{}) {
		utils.Error(w, http.StatusBadRequest, "Resource key not found")
		return
	}
	isRequestValid, errMessage := services.RequestValidator(config.Configs[resourceKey].RequestParams, requestPayload)
	if !isRequestValid {
		utils.Error(w, http.StatusBadRequest, errMessage)
		return
	}

	requestTemplateName := config.Configs[resourceKey].RequestTemplate
	requestTemplateConfig := config.Templates[requestTemplateName]
	method := config.Configs[resourceKey].Method
	queryParams := config.Configs[resourceKey].QueryParams
	url := config.Configs[resourceKey].URL + services.FormQueryParams(method, requestPayload, queryParams)

	requestTemplate := services.FormRequest(requestTemplateConfig, requestPayload)
	out := services.MakeRemoteRequest(url, method, requestTemplate)

	if _, ok := constants.ERROR_MESSAGES[out]; ok {
		utils.Error(w, http.StatusBadRequest, constants.ERROR_MESSAGES[out])
		return
	}
	utils.JSON(w, http.StatusOK, out)
}

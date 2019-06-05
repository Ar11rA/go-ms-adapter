package config

// TemplateParams - JSON params
type TemplateParams struct {
	URL             string `json:"url"`
	RequestTemplate string `json:"requestTemplate"`
	RequestParams   []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"requestParams"`
	ResponseTemplate string `json:"responseTemplate"`
}

// JSONConfig - resource key mapped to template params
type JSONConfig map[string]TemplateParams

// RequestTemplate - resource key mapped to request template
type RequestTemplate map[string][]byte

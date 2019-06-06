package config

// Params - Validation model
type Params struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// TemplateParams - JSON params
type TemplateParams struct {
	URL              string   `json:"url"`
	RequestTemplate  string   `json:"requestTemplate"`
	RequestParams    []Params `json:"requestParams"`
	ResponseTemplate string   `json:"responseTemplate"`
	Method           string   `json:"method"`
}

// JSONConfig - resource key mapped to template params
type JSONConfig map[string]TemplateParams

// RequestTemplate - resource key mapped to request template
type RequestTemplate map[string][]byte

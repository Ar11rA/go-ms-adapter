package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Error http response
func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, map[string]string{"error": message})
}

// JSON http response
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ParseRequest - parse http request
func ParseRequest(r *http.Request) (map[string]interface{}, string) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(string(bodyBuffer)), &result)
	resourceKey := result["resourceKey"].(string)
	delete(result, "resourceKey")
	return result, resourceKey
}

package services

import (
	"bytes"
	"go-ms-adapter/config"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
	"github.com/gojektech/heimdall/httpclient"
)

// RequestValidator - validate the request params
func RequestValidator(requestParams []config.Params, input map[string]interface{}) (bool, string) {
	for _, param := range requestParams {
		if _, ok := input[param.Name]; !ok {
			return false, "Missing input: " + param.Name
		}
		if reflect.TypeOf(input[param.Name]).String() != strings.ToLower(param.Type) {
			return false, "Request failed for " + param.Name
		}
	}
	return true, ""
}

// FormRequest - form the request template based on contents and received data
func FormRequest(contents []byte, d interface{}) *bytes.Buffer {
	t := template.New("request_template")
	s := string(contents)
	t, _ = t.Parse(s)
	buf := &bytes.Buffer{}
	err := t.Execute(buf, d)
	if err != nil {
		log.Fatalf("templating failed with '%s'\n", err)
	}
	return buf
}

// MakeRemoteRequest - make remote call to remote url and input request
func MakeRemoteRequest(remoteURL string, method string, buf *bytes.Buffer) string {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	buff := []io.Reader{buf}
	combined := io.MultiReader(buff...)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	var response *http.Response
	switch(method){
	case "GET": 
		response, _ = client.Get(remoteURL, headers)
		break
	case "POST": 
		response, _ = client.Post(remoteURL, combined, headers)
		break
	case "PUT": 
		response, _ = client.Put(remoteURL, combined, headers)
		break
	case "DELETE": 
		response, _ = client.Delete(remoteURL, headers)
		break
	default : 
		return "INVALID_METHOD"
	}

	log.Println(response.Body)
	byteResponse, _ := ioutil.ReadAll(response.Body)
	return string(byteResponse)
}

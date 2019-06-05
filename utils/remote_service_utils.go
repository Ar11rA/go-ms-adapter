package utils

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

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

// MakePostRequest - make post call to remote url and input request
func MakePostRequest(remoteURL string, buf *bytes.Buffer, out chan string) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	buff := []io.Reader{buf}
	combined := io.MultiReader(buff...)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	response, _ := client.Post(remoteURL, combined, headers)
	byteResponse, _ := ioutil.ReadAll(response.Body)
	out <- string(byteResponse)
}

package unicycle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type FetchOptions struct {
	Method  string
	Query   map[string]string
	Headers map[string]string
	Body    io.Reader
}

func FetchJson[OUTPUT_TYPE any](raw_url string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	true_url, err := url.Parse(raw_url)
	if err != nil {
		return output, err
	}
	if len(options.Query) > 0 {
		query := true_url.Query()
		for key, value := range options.Query {
			query.Set(key, value)
		}
		true_url.RawQuery = query.Encode()
	}
	if options.Method == "" {
		options.Method = "GET"
	}
	request, err := http.NewRequest(options.Method, true_url.String(), options.Body)
	if err != nil {
		return output, err
	}
	for key, value := range options.Headers {
		request.Header.Add(key, value)
	}
	client := http.Client{
		Timeout: time.Minute,
	}
	response, err := client.Do(request)
	if err != nil {
		return output, err
	}
	if (response.StatusCode < 200) || (300 <= response.StatusCode) {
		return output, fmt.Errorf("non-2XX response status code: %d", response.StatusCode)
	}
	response_body_bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output, err
	}
	err = json.Unmarshal(response_body_bytes, &output)
	return output, err
}

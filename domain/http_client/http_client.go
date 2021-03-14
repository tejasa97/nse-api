package http_client

import (
	"fmt"
	"net/http"
	urlpkg "net/url"
	"time"
)

type Headers map[string]string
type Params map[string]string

type Request struct {
	req *http.Request
}

var (
	APIClient apiClientInterface = &apiClient{}
	tr                           = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	request Request
	client  = &http.Client{Transport: tr}
)

type apiClientInterface interface {
	NewGetRequest(string, Headers, Params) (*http.Response, *error)
}

type apiClient struct{}

func (request *Request) generateURLParams(url string, params ...Params) {
	url += "?"
	for _, maps := range params {
		for key, value := range maps {
			url += fmt.Sprintf("&%s=%s", key, value)
		}
	}
	fmt.Println("url", url)
	request.req.URL, _ = urlpkg.Parse(url)
}

func (n *apiClient) NewGetRequest(url string, headers Headers, params Params) (*http.Response, *error) {
	// Çreate a request instance
	var err error
	request.req, err = http.NewRequest("GET", url, nil)

	// Assign the headers and query params (if any)
	if headers != nil {
		for key, value := range headers {
			request.req.Header.Set(key, value)
		}
	}
	if params != nil {
		request.generateURLParams(url, params)
	}
	response, err := client.Do(request.req)
	if err != nil {
		return nil, &err
	}

	return response, nil
}

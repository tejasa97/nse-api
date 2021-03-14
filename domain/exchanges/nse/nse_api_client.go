package nse

import (
	"io/ioutil"

	"github.com/tejasa97/stock-india/domain/http_client"
)

var (
	baseURL string = "https://www1.nseindia.com/"
	headers        = http_client.Headers{
		"Host":             "www1.nseindia.com",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
		"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Connection":       "keep-alive",
		"X-Requested-With": "XMLHttpRequest",
	}
)

type NseApiClient struct {
}

func (n *NseApiClient) newGetRequest(url string, extraHeaders http_client.Headers, params http_client.Params) (string, *error) {

	apiURL := baseURL + url

	// Prepare the headers
	if extraHeaders != nil {
		for key, value := range extraHeaders {
			headers[key] = value
		}
	}
	response, err := http_client.APIClient.NewGetRequest(apiURL, headers, params)
	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(response.Body)

	return string(body), nil
}

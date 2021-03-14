package nse

import (
	"github.com/tejasa97/stock-india/domain/http_client"
)

var (
	NseApiInterface nseApiInterface = &nseApi{}
	nseApiClient                    = NseApiClient{}
)

type nseApiInterface interface {
	GetMarketStatus() (string, *error)
	GetIndices() (string, *error)
	GetSectorsList() (string, *error)
	GetQuoteInfo(string) (string, *error)
	GetNiftyGainers() (string, *error)
	GetNiftyLosers() (string, *error)
	GetAdvancesDeclines() (string, *error)
	GetIndexStocks(string) (string, *error)
}

type nseApi struct {
}

func (n *nseApi) GetMarketStatus() (string, *error) {

	apiEndpoint := "/emerge/homepage/smeNormalMktStatus.json"
	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetIndices() (string, *error) {

	apiEndpoint := "live_market/dynaContent/live_watch/stock_watch/liveIndexWatchData.json"
	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetSectorsList() (string, *error) {

	apiEndpoint := "homepage/peDetails.json"
	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetQuoteInfo(symbol string) (string, *error) {

	apiEndpoint := "live_market/dynaContent/live_watch/get_quote/ajaxGetQuoteJSON.jsp"
	headers := http_client.Headers{
		"Referer": "https://www1.nseindia.com/live_market/dynaContent/live_watch/get_quote/GetQuote.jsp?symbol=" + symbol,
	}
	queryParams := http_client.Params{
		"series": "EQ",
		"symbol": symbol,
	}

	response, err := nseApiClient.newGetRequest(apiEndpoint, headers, queryParams)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetNiftyGainers() (string, *error) {

	apiEndpoint := "live_market/dynaContent/live_analysis/gainers/niftyGainers1.json"

	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetNiftyLosers() (string, *error) {

	apiEndpoint := "live_market/dynaContent/live_analysis/losers/niftyLosers1.json"

	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetAdvancesDeclines() (string, *error) {

	apiEndpoint := "common/json/indicesAdvanceDeclines.json"

	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseApi) GetIndexStocks(symbol string) (string, *error) {

	apiEndpoint := "live_market/dynaContent/live_watch/stock_watch/" + symbol + "StockWatch.json"

	response, err := nseApiClient.newGetRequest(apiEndpoint, nil, nil)
	if err != nil {
		return "", err
	}

	return response, nil
}

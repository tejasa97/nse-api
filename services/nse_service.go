package services

import "github.com/tejasa97/stock-india/domain/exchanges/nse"

var (
	NseService nseServiceInterface = &nseService{}
)

type nseServiceInterface interface {
	GetMarketStatus() (string, *error)
	GetIndices() (string, *error)
	GetSectorsList() (string, *error)
	GetQuoteInfo(string) (string, *error)
	GetNiftyGainers() (string, *error)
	GetNiftyLosers() (string, *error)
	GetAdvancesDeclines() (string, *error)
	GetIndexStocks(string) (string, *error)
}

type nseService struct {
}

func (n *nseService) GetMarketStatus() (string, *error) {
	response, err := nse.NseApiInterface.GetMarketStatus()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetIndices() (string, *error) {
	response, err := nse.NseApiInterface.GetIndices()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetSectorsList() (string, *error) {
	response, err := nse.NseApiInterface.GetSectorsList()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetQuoteInfo(symbol string) (string, *error) {
	response, err := nse.NseApiInterface.GetQuoteInfo(symbol)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetNiftyGainers() (string, *error) {
	response, err := nse.NseApiInterface.GetNiftyGainers()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetNiftyLosers() (string, *error) {
	response, err := nse.NseApiInterface.GetNiftyLosers()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetAdvancesDeclines() (string, *error) {
	response, err := nse.NseApiInterface.GetAdvancesDeclines()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (n *nseService) GetIndexStocks(symbol string) (string, *error) {
	response, err := nse.NseApiInterface.GetIndexStocks(symbol)
	if err != nil {
		return "", err
	}

	return response, nil
}

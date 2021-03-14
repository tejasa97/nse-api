package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tejasa97/stock-india/services"
)

var (
	NseController nseControllerInterface = &nseController{}
)

type nseControllerInterface interface {
	GetMarketStatus(*gin.Context)
	GetIndices(*gin.Context)
	GetSectorsList(*gin.Context)
	GetQuoteInfo(*gin.Context)
	GetNiftyGainers(*gin.Context)
	GetNiftyLosers(*gin.Context)
	GetAdvancesDeclines(*gin.Context)
	GetIndexStocks(*gin.Context)
}

type nseController struct{}

func (n *nseController) GetMarketStatus(c *gin.Context) {
	response, err := services.NseService.GetMarketStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetIndices(c *gin.Context) {
	response, err := services.NseService.GetIndices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetSectorsList(c *gin.Context) {
	response, err := services.NseService.GetSectorsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetQuoteInfo(c *gin.Context) {
	symbol := c.Param("symbol")

	response, err := services.NseService.GetQuoteInfo(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetNiftyGainers(c *gin.Context) {
	response, err := services.NseService.GetNiftyGainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetNiftyLosers(c *gin.Context) {
	response, err := services.NseService.GetNiftyLosers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetAdvancesDeclines(c *gin.Context) {
	response, err := services.NseService.GetAdvancesDeclines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

func (n *nseController) GetIndexStocks(c *gin.Context) {
	symbol := c.Param("symbol")

	response, err := services.NseService.GetIndexStocks(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var bird interface{}
	json.Unmarshal([]byte(response), &bird)

	c.JSON(http.StatusOK, bird)
	return
}

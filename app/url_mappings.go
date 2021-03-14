package app

import (
	"github.com/tejasa97/stock-india/controllers"
)

func mapUrls() {
	// Ping
	router.GET("/nse/market_status", controllers.NseController.GetMarketStatus)

	// NSE
	router.GET("/nse/indices", controllers.NseController.GetIndices)
	router.GET("/nse/sectors_list", controllers.NseController.GetSectorsList)
	router.GET("/nse/quote_info/:symbol", controllers.NseController.GetQuoteInfo)
	router.GET("/nse/nifty_gainers", controllers.NseController.GetNiftyGainers)
	router.GET("/nse/nifty_losers", controllers.NseController.GetNiftyLosers)
	router.GET("/nse/advances_declines", controllers.NseController.GetAdvancesDeclines)
	router.GET("/nse/index_stocks/:symbol", controllers.NseController.GetIndexStocks)
}

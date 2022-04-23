package testhandlers

import "market_apis/functional_test/testdao"

var (
	marketDBConnection = testdao.GetMarketDB().GetConnection()
)

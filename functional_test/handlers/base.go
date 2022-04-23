package handlers

import "market_apis/functional_test/dao"

var (
	marketDBConnection = dao.GetMarketDB().GetConnection()
)

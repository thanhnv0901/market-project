package testmodels

import (
	"market_apis/functional_test/testconfigs"
	"market_apis/functional_test/testdao"
)

var (
	marketDB      = testdao.GetMarketDB()
	configuration = testconfigs.GetConfig()
)

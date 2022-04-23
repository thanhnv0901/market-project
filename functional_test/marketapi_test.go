package functionaltest

import (
	"context"
	"fmt"
	"market_apis/functional_test/handlers"
	"market_apis/functional_test/models"
	"market_apis/functional_test/utils"
	"strconv"

	"github.com/cucumber/godog"
)

func iPrepareRequestWithPayload(payload *godog.DocString) error {
	request := models.GetRequest()
	request.SetPayload(payload.Content)
	return nil
}

func iSendRequestTo(method string, endpoint string) error {

	request := models.GetRequest()
	request.SetMethod(method)

	request.SetURL(endpoint)
	err := request.SendRequest()
	if err != nil {
		return fmt.Errorf("Cannot send request to endpoint %s: %s", endpoint, err.Error())
	}
	return nil
}

func iTruncateTheTable(tableName string) error {

	tableModel := handlers.GetTableModel(tableName)
	err := tableModel.TruncateTable()
	if err != nil {
		return fmt.Errorf("Cannot truncate table %s: %s", tableName, err.Error())
	}
	return nil
}

func theResponseCodeShouldBe(expectation int) error {
	request := models.GetRequest()
	if expectation == request.GetStatusCode() {
		return nil
	}
	return fmt.Errorf("The status code is  %d not same with expectation: %d", request.GetStatusCode(), expectation)
}

func theResponseSuccessShouldBeAndMessageShouldBe(expectedSuccess, expectedMessage string) error {
	request := models.GetRequest()

	success := request.GetIsSuccess()
	message := request.GetMessage()

	tmpExpectedSuccess, err := strconv.ParseBool(expectedSuccess)
	if err != nil {
		return fmt.Errorf("The success status is not bool type: %s", err.Error())
	}
	if success != tmpExpectedSuccess {
		return fmt.Errorf("The success status is %t not same with expectation: %t", success, tmpExpectedSuccess)
	}

	if message != expectedMessage {
		return fmt.Errorf("The message response is %s not same with expectation: %s", message, expectedMessage)
	}
	return nil
}

func dataInTableShoubleBe(tableName string, data *godog.Table) error {

	tableModel := handlers.GetTableModel(tableName)
	actualData, err := tableModel.GetAllData()
	if err != nil {
		return fmt.Errorf("Fail when insert into %s table: %s", tableName, err.Error())
	}

	return utils.TestForEqualData(actualData, data)
}

func iSetupTableTableWithData(tableName string, data *godog.Table) error {

	rows := utils.DataTableConvert(data)
	err := handlers.GetTableModel(tableName).InsertData(rows)
	if err != nil {
		return fmt.Errorf("Fail when insert into %s table: %s", tableName, err.Error())
	}
	return nil
}

func dataWasResponsedIs(data *godog.Table) error {
	request := models.GetRequest()
	actual, err := request.GetDataResponseInMapFormat()
	if err != nil {
		return fmt.Errorf("Fail when parse response data to map: %s", err.Error())
	}

	expectation := utils.DataTableConvert(data)

	return utils.MatchMapData(expectation, actual)
}

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	sc.BeforeSuite(func() {})
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		models.ResetRequest()
		return ctx, nil
	})

	ctx.Step(`^I prepare request with payload:$`, iPrepareRequestWithPayload)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	ctx.Step(`^I truncate the "([^"]*)" table$`, iTruncateTheTable)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response success should be "([^"]*)" and message should be "([^"]*)"$`, theResponseSuccessShouldBeAndMessageShouldBe)
	ctx.Step(`^Data in "([^"]*)" table shouble be:$`, dataInTableShoubleBe)
	ctx.Step(`^I setup table "([^"]*)" table with data:$`, iSetupTableTableWithData)
	ctx.Step(`^the response data should contain$`, dataWasResponsedIs)
}

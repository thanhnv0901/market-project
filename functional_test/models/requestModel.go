package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"market_apis/functional_test/utils"
)

var (
	requestClient *resty.Client
	globalRequest *RequestGlobalModel
	host          string
)

// RequestGlobalModel ..
type RequestGlobalModel struct {
	Method       string
	URL          string
	Payload      string
	Headers      map[string]string
	request      *resty.Request
	response     *resty.Response
	ResponseBody ResponseBodyFromAPI
}

func init() {
	host = "http://localhost:3000/api" // notice

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	requestClient = resty.New().
		SetBaseURL(host).
		SetHeaders(header).
		SetRetryCount(5).
		SetRetryWaitTime(5 * time.Second).SetPreRequestHook(func(c *resty.Client, r *http.Request) error {

		// //Uncommand to debug
		// command, _ := http2curl.GetCurlCommand(r)
		// fmt.Println(command)

		return nil
	})

}

// NewGlobelRequest ..
func NewGlobelRequest() *RequestGlobalModel {
	return &RequestGlobalModel{
		request: requestClient.NewRequest(),
	}
}

// GetRequest ..
func GetRequest() *RequestGlobalModel {
	return globalRequest
}

// SetMethod ..
func (r *RequestGlobalModel) SetMethod(method string) {
	r.Method = method
}

// SetPayload ..
func (r *RequestGlobalModel) SetPayload(payload string) {
	r.Payload = payload
}

// SetURL ..
func (r *RequestGlobalModel) SetURL(endpoint string) error {

	r.URL = endpoint
	return nil
}

// SendRequest ..
func (r *RequestGlobalModel) SendRequest() (err error) {

	switch r.Method {
	case http.MethodGet:
		err = r.SendGetRequest()
	case http.MethodPost:
		err = r.SendPostRequest()
	default:
		return fmt.Errorf("Still not support method: %s", r.Method)
	}

	return err
}

// SendGetRequest ..
func (r *RequestGlobalModel) SendGetRequest() error {

	response, err := r.request.SetResult(&r.ResponseBody).Get(r.URL)
	if err != nil {
		return fmt.Errorf("Error when send request: %s", err.Error())
	}
	r.response = response
	return nil

}

// SendPostRequest ..
func (r *RequestGlobalModel) SendPostRequest() error {

	response, err := r.request.SetBody(r.Payload).SetResult(&r.ResponseBody).Post(r.URL)
	if err != nil {
		return fmt.Errorf("Error when send request: %s", err.Error())
	}

	r.response = response
	return nil
}

// GetStatusCode ..
func (r *RequestGlobalModel) GetStatusCode() int {
	return r.response.StatusCode()
}

// GetIsSuccess ..
func (r *RequestGlobalModel) GetIsSuccess() bool {
	return r.ResponseBody.Success
}

// GetMessage ..
func (r *RequestGlobalModel) GetMessage() string {
	return r.ResponseBody.Message
}

// ResetRequest ..
func ResetRequest() {
	globalRequest = NewGlobelRequest()
}

// GetDataResponseInMapFormat ..
func (r *RequestGlobalModel) GetDataResponseInMapFormat() ([]map[string]interface{}, error) {

	var m []map[string]interface{}

	dataByte, err := json.Marshal(r.ResponseBody.Data)
	if err != nil {
		return nil, fmt.Errorf("Marshall error %s", err)
	}

	m, err = utils.GetResposnseMapForBody(dataByte)

	return m, err
}

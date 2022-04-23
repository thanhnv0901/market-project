package testmodels

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseBodyFromAPI ..
type ResponseBodyFromAPI struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// IngestBody ..
func (r *ResponseBodyFromAPI) IngestBody(res *http.Response) error {

	err := json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return fmt.Errorf("Error when ingest data from response body: %s", err.Error())
	}
	return nil
}

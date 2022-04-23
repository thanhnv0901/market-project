package handlers

// InterfaceTableHandler ..
type InterfaceTableHandler interface {
	TruncateTable() error
	GetAllData() ([]map[string]interface{}, error)
	InsertData([]map[string]interface{}) error
}

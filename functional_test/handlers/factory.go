package handlers

var (
	mapTableModel = make(map[string]InterfaceTableHandler)
)

// Constant table name
const (
	ProductTable = "products"
)

func init() {
	mapTableModel[ProductTable] = &ProductHandler{}
}

// GetTableModel ..
func GetTableModel(table string) InterfaceTableHandler {
	return mapTableModel[table]
}

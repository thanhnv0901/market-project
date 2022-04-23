package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"market_apis/functional_test/utils/errorstrack"
	"reflect"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
)

// ErrorTrackingDefer ..
func ErrorTrackingDefer() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%s", err)
		errTrack := errorstrack.ErrorsTrack{Message: msg}
		log.Println(errTrack.PrintTrackingJSON())
	}
}

// ToJSON ..
func ToJSON(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}

// CompareTwoArray help people can campare two array with any data type
func CompareTwoArray(arrayA interface{}, arrayB interface{}) bool {
	arrayValueA := reflect.ValueOf(arrayA)
	arrayValueB := reflect.ValueOf(arrayB)

	flagChecked := make([]bool, arrayValueA.Len())

	if arrayValueA.Len() != arrayValueB.Len() {
		return false
	}
	for i := 0; i < arrayValueA.Len(); i++ {

		flagExist := false
		for j := 0; j < arrayValueB.Len(); j++ {
			isEqual, _ := CompareTwoObjectInPartly(arrayValueA.Index(i).Interface(), arrayValueB.Index(j).Interface())
			if isEqual && !flagChecked[j] {
				flagExist = true
				flagChecked[j] = true
				break
			}
		}
		if !flagExist {
			return false
		}
	}
	return true
}

// CompareTwoObjectInPartly ..
func CompareTwoObjectInPartly(obj1 interface{}, obj2 interface{}) (bool, error) {

	var (
		valueObject1 = reflect.ValueOf(obj1)
		typeObject1  = reflect.TypeOf(obj1)

		valueObject2 = reflect.ValueOf(obj2)
		typeObject2  = reflect.TypeOf(obj2)
	)

	for i := 0; i < valueObject1.NumField(); i++ {
		f1 := valueObject1.Field(i)

		for j := 0; j < valueObject2.NumField(); j++ {
			f2 := valueObject2.Field(j)

			if typeObject1.Field(i).Name == typeObject2.Field(j).Name {
				if reflect.DeepEqual(f1.Interface(), f2.Interface()) {
					return false, fmt.Errorf(`Field value %s is %v not match with our expectation: %v`, typeObject1.Field(i).Name, f1.Interface(), f2.Interface())
				}
			}
		}
	}

	return true, nil
}

// TestForEqualData ..
func TestForEqualData(actualTable interface{}, expectationTable *godog.Table) error {
	expectationData := DataTableConvert(expectationTable)
	temp, _ := json.Marshal(actualTable)
	var actualData []map[string]interface{}
	json.Unmarshal(temp, &actualData)
	return MatchMapData(expectationData, actualData)
}

//DataTableConvert to convert data table into []map[string]interface{}
func DataTableConvert(dataTable *godog.Table) []map[string]interface{} {
	var headerColumns []string
	for i := 0; i < len(dataTable.Rows[0].Cells); i++ {
		headerColumns = append(headerColumns, strings.TrimSpace(dataTable.Rows[0].Cells[i].Value))
	}
	var dataColumnsArrayFinal [][]interface{}
	for i := 1; i < len(dataTable.Rows); i++ {
		var dataColumnsArray []interface{}
		for j := 0; j < len(dataTable.Rows[i].Cells); j++ {
			dataColumnsArray = append(dataColumnsArray, strings.TrimSpace(dataTable.Rows[i].Cells[j].Value))
		}
		dataColumnsArrayFinal = append(dataColumnsArrayFinal, dataColumnsArray)
	}
	var retMap []map[string]interface{}
	for i := 0; i < len(dataColumnsArrayFinal); i++ {
		var m map[string]interface{}
		m = make(map[string]interface{})
		for j := 0; j < len(headerColumns); j++ {
			m[headerColumns[j]] = dataColumnsArrayFinal[i][j]
		}
		retMap = append(retMap, m)
	}
	return retMap
}

//MatchMapData for api response and data table values
func MatchMapData(dataTableMap []map[string]interface{}, apiReturnMap []map[string]interface{}) error {
	if len(dataTableMap) != len(apiReturnMap) {
		return fmt.Errorf("Different length of data from API request and expected")
	}

	var filterItemsWereCompared []int = make([]int, len(apiReturnMap))
	for i := 0; i < len(dataTableMap); i++ {

		var shouldMoveNextItem bool = false
		for j := 0; j < len(apiReturnMap); j++ {

			if filterItemsWereCompared[j] == 1 {
				continue
			}

			var countFieldCompared int = 0
			for k := range dataTableMap[i] {
				_, ok := apiReturnMap[j][k]
				if !ok {
					return fmt.Errorf("Response does not contain field : %s", k)
				}

				var IsSame bool = true
				switch apiReturnMap[j][k].(type) {
				case float64, float32:
					valAPI := reflect.ValueOf(apiReturnMap[j][k]).Float()
					valData, err := strconv.ParseFloat(reflect.ValueOf(dataTableMap[i][k]).String(), 64)
					if err != nil {
						return fmt.Errorf("Data type mmismatch for column %s and row %d", err, i+1)
					}
					if valData != valAPI {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s ", i+1, k)
						}
						IsSame = false
						break
					}
				case int, int8, int16, int32, int64:
					valAPI := reflect.ValueOf(apiReturnMap[j][k]).Int()
					valData, err := strconv.Atoi(reflect.ValueOf(dataTableMap[i][k]).String())
					if err != nil {
						return fmt.Errorf("Data type mmismatch for column %s and row %d", k, i+1)
					}
					if int64(valData) != valAPI {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s ", i+1, k)
						}
						IsSame = false
						break
					}
				case string:
					valAPI := reflect.ValueOf(apiReturnMap[j][k]).String()
					valData := reflect.ValueOf(dataTableMap[i][k]).String()
					if valData != valAPI {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s ", i+1, k)
						}
						IsSame = false
						break
					}
				case bool:
					valAPI := reflect.ValueOf(apiReturnMap[j][k]).Bool()
					valData, err := strconv.ParseBool(reflect.ValueOf(dataTableMap[i][k]).String())
					if err != nil {
						return fmt.Errorf("Data type mmismatch for column %s and row %d", k, i+1)
					}
					if valData != valAPI {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s", i+1, k)
						}
						IsSame = false
						break
					}
				case nil:
					valData := reflect.ValueOf(dataTableMap[i][k]).String()
					if valData != "" {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s", i+1, k)
						}
						IsSame = false
						break
					}
				case interface{}:
					var dataMap interface{}
					if reflect.ValueOf(dataTableMap[i][k]).String() != "" {
						err := json.Unmarshal([]byte(reflect.ValueOf(dataTableMap[i][k]).String()), &dataMap)
						if err != nil {
							return fmt.Errorf("Error in marshalling data from data table : %s", err)
						}
					}
					valAPI := apiReturnMap[j][k]
					if !reflect.DeepEqual(valAPI, dataMap) {
						if j == len(apiReturnMap)-1 {
							return fmt.Errorf("Mismatch in Data Values for row %d, and column %s", i+1, k)
						}
						IsSame = false
						break
					}
				}

				if !IsSame {
					break
				}

				countFieldCompared++
				if countFieldCompared == len(dataTableMap[i]) {
					shouldMoveNextItem = true
				}
			}

			if shouldMoveNextItem {
				filterItemsWereCompared[j] = 1
				break
			}
		}

	}

	return nil
}

// GetResposnseMapForBody ..
// returns the response from server into []map[string]interface{} form
func GetResposnseMapForBody(responseData []byte) ([]map[string]interface{}, error) {

	var m []map[string]interface{}
	err := json.Unmarshal(responseData, &m)
	if err != nil {
		return nil, fmt.Errorf("Unmarshall error %s", err)
	}

	retMap := make([]map[string]interface{}, 0)
	// sometimes response comes not in the form of array

	val := reflect.ValueOf(m)
	if val.Kind().String() == "slice" {
		for i := 0; i < val.Len(); i++ {
			slicVal := val.Slice(i, i+1)
			mapVal := slicVal.Index(0)
			mapVal2 := mapVal.Interface().(map[string]interface{})
			retMap = append(retMap, mapVal2)
		}
	} else {
		retMap = append(retMap, val.Interface().(map[string]interface{}))
	}
	return retMap, nil
}

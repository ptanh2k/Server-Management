package excel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ReturnedServerData struct {
	Id     uint16 `json:"id,omitempty"`
	Name   string `json:"name"`
	Ip     string `json:"ip"`
	Port   uint16 `json:"port"`
	Status bool   `json:"status"`
}

// Fetch data from API
func FetchData(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return body
}

// s: json data to export to excel
func ExportDataToExcel(data []byte) {

	var jsonData []ReturnedServerData

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	properties := [5]string{"ID", "Name", "IP", "Port", "Status"}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	c1 := 'A'

	asciiValForKey := int(c1)

	var asciiForKey string

	for i := 0; i < len(properties); i++ {
		asciiForKey = string(rune(asciiValForKey))
		xlsx.SetCellValue(sheetName, asciiForKey+"1", properties[i])
		asciiValForKey++
	}

	row := 2

	for i := 0; i < len(jsonData); i++ {
		values := make([]interface{}, 0, len(properties))

		curData := jsonData[i]

		v := reflect.ValueOf(curData)

		for j := 0; j < v.NumField(); j++ {
			values = append(values, v.Field(j).Interface())
		}

		c2 := 'A'

		asciiValForValue := int(c2)

		var asciiForValue string

		for i := 0; i < len(values); i++ {
			asciiForValue = string(rune(asciiValForValue))
			xlsx.SetCellValue(sheetName, asciiForValue+strconv.Itoa(row), values[i])
			asciiValForValue++
		}

		row++
	}

	err = xlsx.SaveAs("../../../result.xlsx")

	if err != nil {
		fmt.Print(err)
		return
	}
}

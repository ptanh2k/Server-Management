package excel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/xuri/excelize/v2"
)

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

	var jsonData interface{}

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	data_servers, _ := jsonData.([]interface{})

	// Set keys
	first := data_servers[0]

	firstColData, _ := first.(map[string]interface{})

	keys := make([]string, 0, len(firstColData))

	for k := range firstColData {
		keys = append(keys, k)
	}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	c1 := 'A'

	asciiValForKey := int(c1)

	var asciiForKey string

	for i := 0; i < len(keys); i++ {
		asciiForKey = string(rune(asciiValForKey))
		xlsx.SetCellValue(sheetName, asciiForKey+"1", keys[i])
		asciiValForKey++
	}

	row := 2

	for _, data_server := range data_servers {
		c2 := 'A'

		asciiValForValue := int(c2)

		var asciiForValue string

		colData, _ := data_server.(map[string]interface{})
		values := make([]interface{}, 0, len(colData)) // Slice of server info
		for _, v := range colData {
			values = append(values, v)
		}

		fmt.Println(values)

		for i := 0; i < len(values); i++ {
			asciiForValue = string(rune(asciiValForValue))
			xlsx.SetCellValue(sheetName, asciiForValue+strconv.Itoa(row), values[i])
			asciiValForValue++
		}
		row++
	}

	err = xlsx.SaveAs("../../result.xlsx")

	if err != nil {
		fmt.Print(err)
		return
	}
}

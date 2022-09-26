package excel

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ImportFromExcel() {
	f, err := excelize.OpenFile("../../../result.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		// Close the spreadsheet
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Slice of Server
	servers := []ReturnedServerData{}

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(row[0])
		name := row[1]
		ip := row[2]
		port, _ := strconv.Atoi(row[3])
		status, _ := strconv.ParseBool(row[4])

		server := ReturnedServerData{uint16(id), name, ip, uint16(port), status}

		servers = append(servers, server)
	}

	jsonServers, _ := json.Marshal(servers)

	fmt.Print(string(jsonServers))
}

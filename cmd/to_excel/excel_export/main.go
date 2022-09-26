package main

import (
	"sm/pkg/excel"
)

func main() {
	res := excel.FetchData("http://localhost:8080/servers")

	excel.ExportDataToExcel(res)
}

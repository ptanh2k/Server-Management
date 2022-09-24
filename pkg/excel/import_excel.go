package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ImportFromExcel() {
	f, err := excelize.OpenFile("../../data/server.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

}

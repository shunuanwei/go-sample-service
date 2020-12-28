package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ImportFromXLS(file string) string {
	fmt.Println(file)
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	// choose your workspace
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			if len(colCell) == 0 {
				fmt.Print("-", "\t")
			}
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	return file
}

func main() {
	ImportFromXLS("E:\\jiexi\\test1.xlsx")
}

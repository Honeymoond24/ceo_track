package data_source

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

type Row struct {
	ID          int    // Not a column
	CompanyBin  string // Column 1 A
	CompanyName string // Column 3 C
	CeoFullName string // Column 22 V
}

type ExcelFile struct {
	FileName string
	Rows     []Row
}

func (e *ExcelFile) ReadFile() ExcelFile {
	f, err := excelize.OpenFile(e.FileName) // Read the file
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read file:", e.FileName)
	startTime := time.Now()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Populate the rows
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Rows:", len(rows))
	for i := 3; i < len(rows); i++ {
		//fmt.Println("Row:", i, len(rows[i]))
		row := Row{ID: i}

		if len(rows[i]) >= 22 {
			row.CeoFullName = rows[i][21]
		} else {
			row.CeoFullName = ""
			fmt.Println("Row:", i, len(rows[i]), "empty")
		}

		if len(rows[i]) >= 3 {
			row.CompanyName = rows[i][2]
		} else {
			fmt.Println("Row:", i, len(rows[i]), "empty")
			continue
		}

		if len(rows[i]) >= 1 {
			row.CompanyBin = rows[i][0]
		} else {
			fmt.Println("Row:", i, len(rows[i]), "empty")
			continue
		}

		e.Rows = append(e.Rows, row)
	}
	fmt.Println(len(e.Rows), "rows read in", time.Since(startTime))
	return *e
}

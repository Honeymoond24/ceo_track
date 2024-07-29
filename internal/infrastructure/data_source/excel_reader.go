package data_source

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"path/filepath"
)

type Row struct {
	ID          int    // Not a column
	CompanyBin  string // Column 1 A
	CompanyName string // Column 3 C
	CeoFullName string // Column 22 V
}

type ExcelFile struct {
	File     *excelize.File
	FileName string
	Rows     []Row
}

//func (e *ExcelFile) ReadFile(fileName string) ExcelFile {
//	e.FileName = filepath.Join("files", fileName)
//	f, err := excelize.OpenFile(e.FileName) // Read the file
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Read file:", e.FileName)
//	startTime := time.Now()
//	defer func() {
//		// Close the spreadsheet.
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//
//	// Populate the rows
//	sheetName := f.GetSheetName(0)
//	rows, err := f.GetRows(sheetName)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Rows:", len(rows))
//	for i := 3; i < len(rows); i++ {
//		//fmt.Println("Row:", i, len(rows[i]))
//		row := Row{ID: i}
//
//		if len(rows[i]) >= 22 {
//			row.CeoFullName = rows[i][21]
//		} else {
//			row.CeoFullName = ""
//			fmt.Println("Row:", i, len(rows[i]), "empty")
//		}
//
//		if len(rows[i]) >= 3 {
//			row.CompanyName = rows[i][2]
//		} else {
//			fmt.Println("Row:", i, len(rows[i]), "empty")
//			continue
//		}
//
//		if len(rows[i]) >= 1 {
//			row.CompanyBin = rows[i][0]
//		} else {
//			fmt.Println("Row:", i, len(rows[i]), "empty")
//			continue
//		}
//
//		e.Rows = append(e.Rows, row)
//	}
//	fmt.Println(len(e.Rows), "rows read in", time.Since(startTime))
//	return *e
//}

// Read reads the file and saves a reference to it
func (e *ExcelFile) Read(fileName string) ExcelFile {
	var err error
	e.FileName = filepath.Join("files", fileName)
	e.File, err = excelize.OpenFile(e.FileName) // Read the file
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read file:", e.FileName)
	defer func() {
		// Close the spreadsheet
		if err := e.File.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return *e
}

// GetSheetRows reads the file and returns an array of iterators for each sheet
func (e *ExcelFile) GetSheetRows() []*excelize.Rows {
	f, err := excelize.OpenFile(e.FileName) // Read the file
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GetRows() Reading file:", e.FileName)
	sheetCount := f.SheetCount
	fmt.Println("Sheet count:", sheetCount)
	var sheetRows []*excelize.Rows
	for i := 0; i < sheetCount; i++ {
		rows, err := f.Rows(f.GetSheetName(i))
		if err != nil {
			fmt.Println(err)
		}
		sheetRows = append(sheetRows, rows)
	}
	return sheetRows
}

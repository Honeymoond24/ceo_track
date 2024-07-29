package application

import (
	"ceo_track/internal/infrastructure/data_source"
	"ceo_track/internal/infrastructure/database"
	"ceo_track/internal/infrastructure/notify"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
)

func createRegions(db *gorm.DB, items []data_source.Item) {
	var regionsName []string
	for _, item := range items {
		regionsName = append(regionsName, item.RegionName)
	}
	regionsRepo := database.RegionRepositoryImpl{}
	regionsRepo.CreateMany(db, regionsName)
}

func processItem(db *gorm.DB, item data_source.Item) {
	excelFile := data_source.ExcelFile{}
	excelFile.Read(item.FileName)
	sheetRows := excelFile.GetSheetRows()

	// Get the region ID
	regionsRepo := database.RegionRepositoryImpl{}
	regionId := regionsRepo.FindByName(db, item.RegionName).ID

	// Create CEOs data in the database
	ceosRepo := database.CeoRepositoryImpl{}
	var ceos []database.Ceo
	var ceoCount int

	for sheetId, rows := range sheetRows {
		var rowId int
		skip := 3
		if sheetId != 0 {
			skip = 1
		}

		for rows.Next() {
			var row data_source.Row
			if rowId < skip {
				rowId++
				continue
			}
			columns, err := rows.Columns()
			if err != nil {
				fmt.Println(err)
			}
			columnsLength := len(columns)
			if columnsLength >= 22 {
				row.CeoFullName = columns[21]
			} else {
				row.CeoFullName = ""
			}
			if columnsLength >= 3 {
				row.CompanyName = columns[2]
			} else {
				continue
			}
			row.CompanyBin = columns[0]

			// Create CEO data every 100 records
			ceos = append(ceos, database.Ceo{
				CompanyBin:  row.CompanyBin,
				CompanyName: row.CompanyName,
				FullName:    row.CeoFullName,
				RegionID:    regionId,
			})
			ceoCount++
			if ceoCount > 100 {
				ceosRepo.CreateMany(db, ceos)
				ceos = nil
				ceoCount = 0
			}
		}
		if err := rows.Close(); err != nil {
			fmt.Println(err)
		}
	}
	ceosRepo.CreateMany(db, ceos)
}

func FirstLaunch(db *gorm.DB) {
	// Get a list of source data files
	dataSource := data_source.DataSource{}
	items := dataSource.GetCeoData()
	fmt.Println("\nitems:", items)

	createRegions(db, items)
	for _, item := range items {
		processItem(db, item)
	}
}

func compareItem(db *gorm.DB, item data_source.Item) ([]CeoNew, []CeoChange) {
	excelFile := data_source.ExcelFile{}
	excelFile.Read(item.FileName)
	sheetRows := excelFile.GetSheetRows()
	var changesNew []CeoNew
	var changesUpdate []CeoChange
	for sheetId, rows := range sheetRows {
		var rowId int
		skip := 3
		if sheetId != 0 {
			skip = 1
		}

		ceosRepo := database.CeoRepositoryImpl{}
		regionsRepo := database.RegionRepositoryImpl{}
		for rows.Next() {
			var ceo database.Ceo
			if rowId < skip {
				rowId++
				continue
			}
			columns, err := rows.Columns()
			if err != nil {
				fmt.Println(err)
			}
			columnsLength := len(columns)
			if columnsLength >= 22 {
				ceo.FullName = columns[21]
			} else {
				ceo.FullName = ""
			}
			if columnsLength >= 3 {
				ceo.CompanyName = columns[2]
			} else {
				continue
			}
			ceo.CompanyBin = columns[0]

			ceoFromDb := ceosRepo.FindByCompanyName(db, ceo.CompanyName)

			if ceoFromDb.ID == 0 { // The CEO is not in the database
				ceo.RegionID = regionsRepo.FindByName(db, item.RegionName).ID
				ceosRepo.CreateMany(db, []database.Ceo{ceo})
				fmt.Println("New CEO data:", ceoFromDb, ceo)
				changesNew = append(changesNew, CeoNew{
					CompanyBin:     ceo.CompanyBin,
					NewCeoFullName: ceo.FullName,
				})
				continue
			}

			if ceoFromDb.CompanyBin == "" {
				continue
			}

			// Check if the data has changed
			if ceoFromDb.CompanyBin == ceo.CompanyBin &&
				ceoFromDb.FullName != ceo.FullName {
				ceosRepo.Update(db, ceoFromDb.ID, ceo)
				fmt.Println("CEO data has changed:", ceoFromDb, ceo)
				changesUpdate = append(changesUpdate, CeoChange{
					CeoNew: CeoNew{
						CompanyBin:     ceo.CompanyBin,
						NewCeoFullName: ceo.FullName,
					},
					OldCeoFullName: ceoFromDb.FullName,
				})
			}
		}
		if err := rows.Close(); err != nil {
			fmt.Println(err)
		}
	}
	return changesNew, changesUpdate
}

func LaunchTrack(db *gorm.DB) {
	// Get a list of source data files
	dataSource := data_source.DataSource{}
	items := dataSource.GetCeoData()
	fmt.Println("\nitems:", items)

	var changes CeoChanges
	for _, item := range items {
		changesNew, changesUpdate := compareItem(db, item)
		changes.New = append(changes.New, changesNew...)
		changes.Changes = append(changes.Changes, changesUpdate...)
	}
	message := changes.MakeMessageReport()
	notifier := notify.Notifier{ChatID: os.Getenv("CHAT_ID"), BotToken: os.Getenv("BOT_TOKEN")}
	err := notifier.SendMessage(message)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
}

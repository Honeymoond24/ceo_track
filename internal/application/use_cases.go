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

func FirstLaunch(db *gorm.DB) {
	// This is the first launch of the application, so we need to prepare the database

	// Get CEO data from the data source
	dataSource := data_source.DataSource{}
	items := dataSource.GetCeoData()
	fmt.Println("\nitems:", items)

	// Create regions
	var regionsName []string
	for _, item := range items {
		regionsName = append(regionsName, item.RegionName)
	}

	regionsRepo := database.RegionRepositoryImpl{}
	regionsRepo.CreateMany(db, regionsName)

	// Create CEOs data in the database
	ceosRepo := database.CeoRepositoryImpl{}
	for _, item := range items {
		excelFile := data_source.ExcelFile{}
		if item.FileName == "" {
			fmt.Println("item.FileName is empty:", item)
			continue
		}
		excelFile.ReadFile(item.FileName)
		fmt.Println("\nExcelFile:", len(excelFile.Rows))

		regionID := regionsRepo.FindByName(db, item.RegionName).ID
		var ceos []database.Ceo
		for _, row := range excelFile.Rows {
			ceos = append(ceos, database.Ceo{
				CompanyBin:  row.CompanyBin,
				CompanyName: row.CompanyName,
				FullName:    row.CeoFullName,
				RegionID:    regionID,
			})
		}
		ceosRepo.CreateMany(db, ceos)
		//break
	}
}

func LaunchTrack(db *gorm.DB) {
	// This is the regular launch of the application
	// Get CEO data from the data source
	dataSource := data_source.DataSource{}
	items := dataSource.GetCeoData()
	fmt.Println("\nitems:", items)

	//var fileCeos [][]database.Ceo
	regionsRepo := database.RegionRepositoryImpl{}

	var changes CeoChanges
	for i, item := range items {
		excelFile := data_source.ExcelFile{}
		excelFile.ReadFile(item.FileName)
		fmt.Println("\n", i, "ExcelFile:", len(excelFile.Rows))
		region := regionsRepo.FindByName(db, item.RegionName)
		regionID := region.ID
		var ceos []database.Ceo
		for _, row := range excelFile.Rows {
			ceos = append(ceos, database.Ceo{
				CompanyBin:  row.CompanyBin,
				CompanyName: row.CompanyName,
				FullName:    row.CeoFullName,
				RegionID:    regionID,
			})
		}
		// Compare the data from the data source with the data in the database
		fmt.Println("fileCeo:", len(ceos))
		for _, ceo := range ceos {
			ceosRepo := database.CeoRepositoryImpl{}
			ceoFromDb := ceosRepo.FindByCompanyName(db, ceo.CompanyName)
			if ceoFromDb.CompanyBin == "" {
				continue
			}

			if ceoFromDb.ID == 0 { // The CEO is not in the database
				ceosRepo.CreateMany(db, []database.Ceo{ceo})
				fmt.Println("New CEO data:", ceoFromDb, ceo)
				changes.New = append(changes.New, CeoNew{
					CompanyBin:     ceo.CompanyBin,
					NewCeoFullName: ceo.FullName,
				})
				continue
			}

			// Check if the data has changed
			if ceoFromDb.CompanyBin == ceo.CompanyBin && ceoFromDb.FullName != ceo.FullName {
				ceosRepo.Update(db, ceoFromDb.ID, ceo)
				fmt.Println("CEO data has changed:", ceoFromDb, ceo)
				changes.Changes = append(changes.Changes, CeoChange{
					CeoNew: CeoNew{
						CompanyBin:     ceo.CompanyBin,
						NewCeoFullName: ceo.FullName,
					},
					OldCeoFullName: ceoFromDb.FullName,
				})
			}
		}
	}
	message := changes.MakeMessageReport()
	notifier := notify.Notifier{ChatID: os.Getenv("CHAT_ID"), BotToken: os.Getenv("BOT_TOKEN")}
	err := notifier.SendMessage(message)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
}

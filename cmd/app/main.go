package main

import (
	"ceo_track/internal/application"
	"ceo_track/internal/infrastructure/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("GORM_DB_URI")
	fmt.Println(dsn)

	// Connect to the database
	db := database.Connection(dsn)
	fmt.Println("DB:", db)

	// Migrate the schema
	err = db.AutoMigrate(&database.Ceo{}, &database.Region{})
	if err != nil {
		log.Fatal("Error migrating database")
	}

	var ceoCount int64
	db.Model(&database.Ceo{}).Count(&ceoCount)
	fmt.Println("Ceo count at the start:", ceoCount)

	if ceoCount == 0 {
		application.FirstLaunch(db)
	} else {
		application.LaunchTrack(db)
	}
	//dataSource := data_source.DataSource{}
	//dataFiles := dataSource.GetCeoData()
	//fmt.Println("DataFiles:", dataFiles)
	//for _, dataFile := range dataFiles {
	//	excelFile := data_source.ExcelFile{FileName: dataFile}
	//	excelFile.ReadFile()
	//}
}

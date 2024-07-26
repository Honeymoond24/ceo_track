package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connection(dsn string) *gorm.DB {
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

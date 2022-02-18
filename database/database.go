package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Couldn't close the database")
	}
	return sqlDB.Close()
}

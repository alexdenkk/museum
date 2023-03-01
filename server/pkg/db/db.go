package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	return db, err
}

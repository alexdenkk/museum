package db

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbName, dbPort, dbUser, dbPassword string) (*gorm.DB, error) {
	dsn := []string{
		"dbname=" + dbName,
		"port=" + dbPort,
		"user=" + dbUser,
		"password=" + dbPassword,
		"host=db",
		"sslmode=disable",
		"TimeZone=Asia/Novosibirsk",
	}

	db, err := gorm.Open(
		postgres.Open(strings.Join(dsn, " ")),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	return db, err
}

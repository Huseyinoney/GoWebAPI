package Database

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(""), &gorm.Config{})
	if err != nil {
		panic("failed to connection database")
	}
	return db
}

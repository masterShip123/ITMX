package utilities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"itmx/types"
)

func SetupDB() *gorm.DB {
	var err error
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto Migrate
	db.AutoMigrate(&types.Customer{})
	return db
}

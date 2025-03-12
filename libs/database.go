package libs

import (
	"blog/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	// init gorm sqlite database
	db, err := gorm.Open(sqlite.Open("data/db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema
	db.AutoMigrate(
		&models.Comment{},
		&models.Post{},
		&models.User{},
	)

	DB = db
}

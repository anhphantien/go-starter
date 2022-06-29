package repositories

import (
	"go-starter/entities"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Sync(db *gorm.DB) {
	db.AutoMigrate(
		entities.Book{},
		entities.User{},
	)
	DB = db
}

func CreateSqlBuilder(model any) *gorm.DB {
	return DB.Model(model)
}

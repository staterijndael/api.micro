package common

import (
	"github.com/deissh/api.micro/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/gommon/log"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "tests.db")
	if err != nil {
		log.Panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

// Migrate all needed tables
func Migrate() {
	// create tables if not exist
	// todo: add auto migration
	DB.AutoMigrate(&models.Token{}, &models.User{})
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

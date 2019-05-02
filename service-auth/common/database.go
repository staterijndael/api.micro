package common

import (
	"github.com/deissh/api.micro/models"
	"github.com/deissh/api.micro/service-auth/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	host := helpers.GetEnv("DB_HOST", "127.0.0.1")
	user := helpers.GetEnv("DB_USER", "postgres")
	dbName := helpers.GetEnv("DB_NAME", "microapi")
	psw := helpers.GetEnv("DB_PSW", "postgres")

	db, err := gorm.Open("postgres", "sslmode=disable host="+host+" user="+user+" dbname="+dbName+" password="+psw)
	if err != nil {
		log.Panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db

	log.Info("Database connected")
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

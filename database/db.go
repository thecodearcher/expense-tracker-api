package database

import (
	"expense-tracker/models"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //import sql
	"github.com/jinzhu/gorm"
)

//DB instance of the database
var DB *gorm.DB

//InitializeDb connects to MySQL db and returns *gorm.DB
func InitializeDb() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHost, dbName)
	db, err := gorm.Open("mysql", sqlStr)

	DB = db
	if err != nil {
		panic("Failed to connect database")
	}
	log.Print("Database connection established!")

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	return db
}

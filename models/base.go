package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB 

func init() {
	e := godotenv.Load() // Load .env file
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable port=%s password=%s", dbHost, username, dbName, dbPort, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}) // Database migration
}

func GetDB() *gorm.DB {
	return db
}
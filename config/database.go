package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func init() {
	// Load env from .env
	godotenv.Load()
	connectDatabase()
}

func connectDatabase() {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
	))
	if err != nil {
		log.Fatalln(err)
	}
}
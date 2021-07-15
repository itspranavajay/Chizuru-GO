package database

import (
	"Chizuru-GO/chizuru/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var SESSION *gorm.DB

func StartDB() {
	log.Printf("Database: %s", config.DBUrl)
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}

	SESSION = db

	log.Println("Database connected ")

	//Create tables if they don't exist
	SESSION.AutoMigrate(&User{}, &Chat{})
	log.Println("Auto-migrated database schema")

}

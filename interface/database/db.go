package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDB() {
	d, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database sucessfully")
	DB = d
}

func AutoMigrate() {
	log.Println("Auto Migrating Models...")
	err := DB.AutoMigrate(&Repository{}, &Commit{}, &User{})
	if err != nil {
		panic(err)
	}
	log.Println("Migrated DB Successfully")
}

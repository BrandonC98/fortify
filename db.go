package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/passMan_db", config.DBUser, config.DBPassword, config.DBHost)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Credentials{})

	return db
}

func AddCredsRecord(creds *Credentials, db *gorm.DB) {
	result := db.Create(creds)
	if result.Error != nil {
		log.Fatal(result.Error)

	}
}

func retriveAllCreds(db *gorm.DB) []Credentials {

	var creds []Credentials
	db.Find(&creds)

	return creds
}

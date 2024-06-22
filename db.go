package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
)

func Database(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", config.DBUser, config.DBPassword, config.DBHost)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	dbName := "passman_db"

	if err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error; err != nil {
		slog.Error(err.Error())
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DBUser, config.DBPassword, config.DBHost, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.AutoMigrate(&Credentials{})
	if err != nil {
		slog.Error(err.Error())
	}

	return db
}

func AddCredsRecord(creds *Credentials, db *gorm.DB) {
	result := db.Create(creds)
	if result.Error != nil {
		slog.Error(result.Error.Error())

	}
}

func retriveAllCreds(db *gorm.DB) []Credentials {

	var creds []Credentials
	db.Find(&creds)

	return creds
}

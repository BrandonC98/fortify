package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
)

type CredentialRepository struct {
	gorm.DB
	user     string
	password string
	host     string
	name     string
}

type CredsRepo interface {
	AddCredsRecord(*Credentials)
	retriveAllCreds() []Credentials
}

func newCredentialRepository(host string, name string, user string, password string) *CredentialRepository {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	r := CredentialRepository{
		DB:   *db,
		user: user,
		host: host,
		name: name, password: password,
	}

	return &r
}

func (r *CredentialRepository) Setup() {
	if err := r.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", r.name)).Error; err != nil {
		slog.Error(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", r.user, r.password, r.host, r.name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	r.DB = *db

	err = r.AutoMigrate(&Credentials{})
	if err != nil {
		slog.Error(err.Error())
	}
}

func (r *CredentialRepository) AddCredsRecord(creds *Credentials) {
	result := r.Create(creds)
	if result.Error != nil {
		slog.Error(result.Error.Error())
	}
}

func (r *CredentialRepository) retriveAllCreds() []Credentials {
	var creds []Credentials
	r.Find(&creds)

	return creds
}

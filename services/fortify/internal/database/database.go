package database

import (
	"fmt"

	"log/slog"

	"github.com/BrandonC98/fortify/services/fortify/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SecretsRepository struct {
	gorm.DB
	user     string
	password string
	host     string
	name     string
}

type Repository interface {
	AddRecord(*model.Secret)
	RetriveAllRecords() []model.Secret
}

func NewSecretsRepository(host string, name string, user string, password string) *SecretsRepository {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error(err.Error())
	}

	r := SecretsRepository{
		DB:       *db,
		user:     user,
		host:     host,
		name:     name,
		password: password,
	}

	return &r
}

func (r *SecretsRepository) Setup() {
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

	err = r.AutoMigrate(&model.Secret{})
	if err != nil {
		slog.Error(err.Error())
	}
}

func (r *SecretsRepository) AddRecord(creds *model.Secret) {
	result := r.Create(creds)
	if result.Error != nil {
		slog.Error(result.Error.Error())
	}
}

func (r *SecretsRepository) RetriveAllRecords() []model.Secret {
	var creds []model.Secret
	r.Find(&creds)

	return creds
}

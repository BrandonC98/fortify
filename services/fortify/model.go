package main

type Credentials struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Passwd string
}

type Config struct {
	Port       int
	PassGenURL string
	DBUser     string
	DBHost     string
	DBPassword string
}

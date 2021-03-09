package database

import (
	"demo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	dns := "host=localhost user=sirawichvongchuy password=1w2e3r4t5y dbname=auther"
	connection, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Cloud not connect to database ⚠️")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}

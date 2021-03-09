package database

import (
	"demo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func envFetcher(key string) string {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalf("Error Cannot load env file")
	}
	return os.Getenv(key)
}

var DB *gorm.DB

func Connect() {
	dns := "host=" + envFetcher("DB_HOST") + " " + "user=" + envFetcher("DB_USERNAME") + " " + "password=" + envFetcher("DB_PASSWORD") + " " + "dbname=" + envFetcher("DB_NAME")
	connection, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Cloud not connect to database ⚠️")
	}
	DB = connection
	if err := connection.AutoMigrate(&models.User{}); err != nil {
		panic("Cannot migrate to user table")
	}
}

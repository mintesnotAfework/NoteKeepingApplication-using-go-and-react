package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can not fine the .env file")
	}

	url := os.Getenv("POSTGRES_URL")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	DBName := os.Getenv("POSTGRES_DATABASE_NAME")
	pgSsl := os.Getenv("POSTGERS_SSL")

	dsn := "host=" + url + " user=" + username + " password=" + password + " dbname=" + DBName + " port=" + port + " sslmode=" + pgSsl + " TimeZone=UTC"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

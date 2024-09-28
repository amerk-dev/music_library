package store

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"test/internal/models"

	"log"
)

var Db *gorm.DB

func DataBaseInit() {
	host := os.Getenv("POSTGRES_HOST")
	port := 5432
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = Db.AutoMigrate(&models.Song{})
	if err != nil {
		log.Println(err)
	}
}

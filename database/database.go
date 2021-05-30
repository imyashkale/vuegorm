package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=p@ssw0rd dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	return db
}

package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectPostgresDB(dsn string) (db *gorm.DB) {
	var err error

	//dsn := "host=localhost user=postgres password=postgres dbname=video_annotation port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "ERROR DB CONNECTION")
	}

	return

}

package database

import (
	"evolve/models"
	"log"
)

func Migrate() {
	// Migrate the schema
	db := DB
	//ml := db.Config
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatalln(err)
	}
	seed(db)
}

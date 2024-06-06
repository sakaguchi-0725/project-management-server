package testutil

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"todo",
		"54321",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("test database connection error: %v", err)
	}

	return db
}

func FlushRecords(db *gorm.DB, mdl interface{}) {
	db.Where("1 = 1").Delete(mdl)
}

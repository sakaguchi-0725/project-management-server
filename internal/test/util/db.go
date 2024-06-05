package testutil

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() *gorm.DB {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_PORT"),
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

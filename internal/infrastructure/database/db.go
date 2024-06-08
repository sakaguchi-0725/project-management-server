package database

import (
	"fmt"
	"log"

	"github.com/sakaguchi-0725/go-todo/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB.Host,
		config.DB.User,
		config.DB.Passwrod,
		config.DB.Name,
		config.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	return db
}

package testutil

import (
	"fmt"

	"github.com/sakaguchi-0725/go-todo/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.TestDB.Host,
		config.TestDB.User,
		config.TestDB.Passwrod,
		config.TestDB.Name,
		config.TestDB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FlushRecords(db *gorm.DB, mdl interface{}) {
	db.Where("1 = 1").Delete(mdl)
}

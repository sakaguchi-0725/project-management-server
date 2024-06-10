package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Name     string
	Port     string
	User     string
	Passwrod string
}

type TestDBConfig struct {
	Host     string
	Name     string
	Port     string
	User     string
	Passwrod string
}

var (
	DB     DBConfig
	TestDB TestDBConfig
)

func LoadDBConfig() {
	DB.Host = os.Getenv("DB_HOST")
	DB.Name = os.Getenv("DB_NAME")
	DB.Port = os.Getenv("DB_PORT")
	DB.User = os.Getenv("DB_USER")
	DB.Passwrod = os.Getenv("DB_PASSWORD")
}

func LoadTestDBConfig() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Printf("環境変数の読み込みに失敗しました: %v\n", err)
	}

	TestDB.Host = os.Getenv("TEST_DB_HOST")
	TestDB.Name = os.Getenv("TEST_DB_NAME")
	TestDB.Port = os.Getenv("TEST_DB_PORT")
	TestDB.User = os.Getenv("TEST_DB_USER")
	TestDB.Passwrod = os.Getenv("TEST_DB_PASSWORD")
}

package config

import "os"

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
	TestDB.Host = os.Getenv("TEST_DB_HOST")
	TestDB.Name = os.Getenv("TEST_DB_NAME")
	TestDB.Port = os.Getenv("TEST_DB_PORT")
	TestDB.User = os.Getenv("TEST_DB_USER")
	TestDB.Passwrod = os.Getenv("TEST_DB_PASSWORD")
}

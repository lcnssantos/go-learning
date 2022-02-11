package shared

import "os"

type Configuration struct {
	DB_DRIVER    string
	DB_HOST      string
	DB_USER      string
	DB_PASS      string
	DB_PORT      string
	DB_NAME      string
	DB_POOL_SIZE string
}

func GetEnvironmentConfiguration() *Configuration {
	return &Configuration{
		DB_DRIVER:    os.Getenv("DB_DRIVER"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASS:      os.Getenv("DB_PASS"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_NAME:      os.Getenv("DB_NAME"),
		DB_POOL_SIZE: os.Getenv("DB_POOL_SIZE"),
	}
}

package shared

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

func Connect() (*sql.DB, error) {
	configuration := GetEnvironmentConfiguration()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USER, configuration.DB_PASS, configuration.DB_NAME)

	db, err := sql.Open(configuration.DB_DRIVER, connectionString)

	if err != nil {
		return &sql.DB{}, err
	}

	poolSize, _ := strconv.Atoi(configuration.DB_POOL_SIZE)

	db.SetMaxOpenConns(poolSize)

	return db, nil
}

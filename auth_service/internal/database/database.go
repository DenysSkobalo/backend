package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Connect(config Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Connected to database")

	return db, nil
}

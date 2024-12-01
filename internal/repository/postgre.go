package repository

import (
	"fmt"
	"log"
	"urlshortener/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(config *config.Config) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SSLMode,
	)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return

}

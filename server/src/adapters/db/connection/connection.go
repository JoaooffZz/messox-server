package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type ConfigDB struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
}

func New() *ConfigDB {
	return &ConfigDB{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_NAME"),
	}
}

func GetConn(config *ConfigDB) (*sql.DB, error) {
	urlDB := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.USER,
		config.PASSWORD,
		config.HOST,
		config.PORT,
		config.DBNAME,
	)
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping() 
	if err != nil {
		return nil, err
    }

	return db, nil
}
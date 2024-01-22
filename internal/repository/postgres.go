package repository

import (
	"database/sql"
	"fmt"
	"test_task/config"

	_ "github.com/lib/pq"
)

func InitPostgresDB(cfg config.ConfigDB) (*sql.DB, error) {
	connData := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode= %s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sql.Open(cfg.Driver, connData)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConnector struct {
	DB *sqlx.DB
}

func New(cfg *DBConnConfig) (*DBConnector, error) {
	sqlConfig := mysql.Config{
		User:   cfg.Username,
		Passwd: cfg.Password,
		Addr:   cfg.Host,
		Net:    "tcp",
		DBName: cfg.Database,
	}

	conn, err := sqlx.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &DBConnector{
		DB: conn,
	}, nil
}

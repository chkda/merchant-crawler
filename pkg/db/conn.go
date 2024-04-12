package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connector struct {
	DB *sqlx.DB
}

func New(cfg *DBConnConfig) (*Connector, error) {
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

	return &Connector{
		DB: conn,
	}, nil
}

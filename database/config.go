package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type DBManager struct {
	dbPool *pgxpool.Pool
}

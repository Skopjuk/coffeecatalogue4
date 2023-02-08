package repository

import (
	"coffeecatalogue4/pkg/logging"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable      = "users"
	coffeesTable    = "coffees"
	roasteriesTable = "coffees"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	logger := logging.GetLogger()

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Error("file didn't open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("db do not response")
		return nil, err
	}
	logger.Info("new db created")
	return db, nil
}

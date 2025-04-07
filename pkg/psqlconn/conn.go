package psqlconn

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	driverName = "pgx"
)

func getConnString(cfg Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func CreateConnection(cfg Config) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect(driverName, getConnString(cfg))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Stop(db *sqlx.DB) error {
	return db.Close()
}

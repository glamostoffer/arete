package psqlconn

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Connector struct {
	cfg Config
	db  *sqlx.DB
}

const (
	driverName = "pgx"
)

func (c *Connector) getConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.cfg.Host,
		c.cfg.Port,
		c.cfg.User,
		c.cfg.Password,
		c.cfg.DBName,
		c.cfg.SSLMode,
	)
}

func (c *Connector) CreateConnection() (db *sqlx.DB, err error) {
	c.db, err = sqlx.Connect(driverName, c.getConnString())
	if err != nil {
		return nil, err
	}

	err = c.db.Ping()
	if err != nil {
		return nil, err
	}

	return c.db, nil
}

func (c *Connector) Stop() error {
	return c.db.Close()
}

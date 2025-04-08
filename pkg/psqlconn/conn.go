package psqlconn

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driverName    = "postgres"
	componentName = "psqlconn"
)

type Connector struct {
	*sqlx.DB
	cfg Config
}

func New(cfg Config) Connector {
	conn := Connector{
		cfg: cfg,
	}

	db, err := sqlx.Connect(driverName, getConnString(cfg))
	if err != nil {
		log.Fatal(err)
	}

	conn.DB = db

	return conn
}

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

func (c *Connector) Start(_ context.Context) (err error) {
	err = c.DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	return c.DB.Close()
}

func (c *Connector) GetName() string {
	return componentName
}

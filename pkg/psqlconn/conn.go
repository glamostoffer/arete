package psqlconn

import (
	"context"
	"fmt"

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
	return Connector{
		cfg: cfg,
	}
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
	c.DB, err = sqlx.Connect(driverName, getConnString(c.cfg))
	if err != nil {
		return err
	}

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

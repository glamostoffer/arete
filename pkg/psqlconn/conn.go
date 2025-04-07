package psqlconn

import (
	"context"
	"fmt"

	"github.com/glamostoffer/arete/pkg/component"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driverName    = "pgx"
	componentName = "psqlconn"
)

type connector struct {
	db  *sqlx.DB
	cfg Config
}

func New(cfg Config) component.Component {
	return &connector{
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

func (c *connector) Start(ctx context.Context) (err error) {
	c.db, err = sqlx.Connect(driverName, getConnString(c.cfg))
	if err != nil {
		return err
	}

	err = c.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (c *connector) Stop(ctx context.Context) error {
	return c.db.Close()
}

func (c *connector) GetName() string {
	return componentName
}

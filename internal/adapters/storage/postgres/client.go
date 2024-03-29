package postgres

import (
	"context"
	"fmt"
	"os"
	"skillbox/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

//"postgres://postgres:Wild54323@localhost:5432/postgres"
type Client struct {
	ctx        context.Context
	ConnUrl    string
	l          *logrus.Logger
	Connection *pgx.Conn
}

func NewClient(logger *logrus.Logger, cfg *config.Config) *Client {
	return &Client{
		ctx:     context.Background(),
		ConnUrl: cfg.PG.URL,
		l:       logger,
	}
}

//Connect...
func (c *Client) Connect() error {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		c.l.Fatalf("migrate: environment variable not declared: PG_URL")
	}
	//testurl := "postgres://postgres:Wild54323@localhost:5432/postgres"
	databaseURL += "?sslmode=disable"
	fmt.Println("URL:", c.ConnUrl)
	conn, err := pgx.Connect(c.ctx, databaseURL)
	if err != nil {
		return err
	}
	c.Connection = conn
	return nil
}

//Close...
func (c *Client) Close() error {
	return c.Connection.Close(c.ctx)
}

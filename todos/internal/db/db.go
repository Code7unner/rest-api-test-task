package db

import (
	"net"
	"strconv"

	"github.com/code7unner/rest-api-test-task/todos/config"
	"github.com/go-pg/pg/v10"
)

// Connect ...
func Connect(cfg *config.Config) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.DbName,
	})
	if err := db.Ping(db.Context()); err != nil {
		return nil, err
	}
	return db, nil
}

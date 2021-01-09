package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/labstack/gommon/log"
)

func init() {
	migrations.MustRegister(createUsers, rollbackUsers)
}

func createUsers(db migrations.DB) error {
	log.Info("creating table [users]...")
	_, err := db.Exec(
		`CREATE TABLE users (
			id bigserial NOT NULL primary key,
			username varchar UNIQUE NOT NULL,
			password varchar NOT NULL
		);		
	`)

	return err
}

func rollbackUsers(db migrations.DB) error {
	log.Warn("dropping table [users]...")
	_, err := db.Exec(`DROP TABLE users`)

	return err
}

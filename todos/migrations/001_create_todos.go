package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/labstack/gommon/log"
)

func init() {
	migrations.MustRegister(createTodos, rollbackTodos)
}

func createTodos(db migrations.DB) error {
	log.Info("creating table [todos]...")
	_, err := db.Exec(
		`CREATE TABLE todos (
			id bigserial NOT NULL primary key,
			user_id bigserial NOT NULL,
			title varchar NOT NULL,
			description varchar NOT NULL,
			time_to_complete timestamp not null,
			constraint "user_id_fk" foreign key ("user_id") references "users"("id")
		);		
	`)

	return err
}

func rollbackTodos(db migrations.DB) error {
	log.Warn("dropping table [todos]...")
	_, err := db.Exec(`DROP TABLE todos`)

	return err
}

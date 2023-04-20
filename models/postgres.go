package models

import (
	"fmt"

	"github.com/bonhokage06/lenslocked/database"
	"github.com/bonhokage06/lenslocked/migrations"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
}

// up migration
func (p *Postgres) Migrate() error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %v", err)
	}
	err = goose.Up(database.Db.DB(), ".")
	if err != nil {
		return fmt.Errorf("migrate: %v", err)
	}
	return nil
}
func (p *Postgres) MigrateFs() error {
	goose.SetBaseFS(migrations.FS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	p.Migrate()
	return nil
}

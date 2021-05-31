package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

//go:embed *.go
var goMigrations embed.FS

func init() {
	if err := Migrations.Discover(goMigrations); err != nil {
		panic(err)
	}
}

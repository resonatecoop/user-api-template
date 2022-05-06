package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/resonatecoop/user-api/model"
)

func init() {

	// Drop and create tables.
	models := []interface{}{
		(*model.UploadSubmission)(nil),
		(*model.Track)(nil),
		(*model.TrackGroup)(nil),
		(*model.Play)(nil),
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		if _, err := db.Exec(`CREATE TYPE track_status AS ENUM ('paid', 'free', 'both');`); err != nil {
			return err
		}

		if _, err := db.Exec(`CREATE TYPE play_type AS ENUM ('paid', 'free');`); err != nil {
			return err
		}

		if _, err := db.Exec(`CREATE TYPE track_group_type AS ENUM ('lp', 'ep', 'single', 'playlist');`); err != nil {
			return err
		}

		for _, model := range models {
			_, err := db.NewDropTable().Model(model).IfExists().Exec(ctx)
			if err != nil {
				panic(err)
			}

			_, err = db.NewCreateTable().Model(model).Exec(ctx)
			if err != nil {
				panic(err)
			}
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		for _, this_model := range models {
			_, err := db.NewDropTable().Model(this_model).IfExists().Exec(ctx)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
}

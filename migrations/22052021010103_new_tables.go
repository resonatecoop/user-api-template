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
		(*model.Credit)(nil),
		(*model.UserMembership)(nil),
		(*model.StripeUser)(nil),
		(*model.ShareTransaction)(nil),
		(*model.MembershipClass)(nil),
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

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

package migrations

import (
	//	"github.com/go-pg/migrations"
	//	"github.com/go-pg/pg/orm"
	"context"
	"fmt"

	"github.com/uptrace/bun"

	//"github.com/uptrace/bun/dialect/pgdialect"
	//_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/resonatecoop/user-api-template/model"
)

func init() {

	// Drop and create tables.
	models := []interface{}{
		(*model.UserGroup)(nil),
		(*model.StreetAddress)(nil),
		(*model.Tag)(nil),
		(*model.Role)(nil),
		(*model.User)(nil),
		(*model.Link)(nil),
		//		(*model.UserGroupPrivacy)(nil),
		(*model.GroupType)(nil),
		//		(*model.UserGroupMember)(nil),
		(*model.EmailToken)(nil),
		(*model.Client)(nil),
		(*model.Scope)(nil),
		(*model.RefreshToken)(nil),
		(*model.AuthorizationCode)(nil),
		(*model.AccessToken)(nil),
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		// if _, err := db.QueryContext(ctx, `CREATE EXTENSION IF NOT EXISTS "hstore"`); err != nil {
		// 	return err
		// }

		// if _, err := db.QueryContext(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		// 	return err
		// }

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

	// Migrations.MustRegister(func(db migrations.DB) error {
	// if _, err := db.Exec( /* language=sql */ `CREATE EXTENSION IF NOT EXISTS "hstore"`); err != nil {
	// 	return err
	// }
	//
	// if _, err := db.Exec( /* language=sql */ `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
	//   return err
	// }
	// if _, err := db.Exec(`CREATE TYPE github.com/resonatecoop/_status AS ENUM ('paid', 'free', 'both');`); err != nil {
	// 	return err
	// }

	// if _, err := db.Exec(`CREATE TYPE play_type AS ENUM ('paid', 'free');`); err != nil {
	// 	return err
	// }

	// if _, err := db.Exec(`CREATE TYPE github.com/resonatecoop/_group_type AS ENUM ('lp', 'ep', 'single', 'playlist');`); err != nil {
	// 	return err
	// }

	// for _, model := range []interface{}{
	// 	&model.StreetAddress{},
	// 	&model.Tag{},
	// 	&model.User{},
	// 	&model.Link{},
	// 	&model.UserGroupPrivacy{},
	// 	&model.GroupTaxonomy{},
	// 	&model.UserGroup{},
	// 	// &model.github.com/resonatecoop/{},
	// 	// &model.github.com/resonatecoop/Group{},
	// 	&model.UserGroupMember{},
	// 	// &model.Play{},
	// } {
	// 	err := orm.CreateTable(db.(orm.DB), model, &orm.CreateTableOptions{
	// 		FKConstraints: true,
	// 		IfNotExists:   true,
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// 	db.RegisterModel((*model.UserGroupMember)(nil))
	// 	if _, err := db.Exec(`alter table user_group_members add foreign key (user_group_id) references user_groups(id)`); err != nil {
	// 		return err
	// 	}
	// 	if _, err := db.Exec(`alter table user_group_members add foreign key (member_id) references user_groups(id)`); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }, func(db migrations.DB) error {
	// 	if _, err := db.Exec(`DROP TYPE IF EXISTS play_type CASCADE;`); err != nil {
	// 		return err
	// 	}
	// 	if _, err := db.Exec(`DROP TYPE IF EXISTS github.com/resonatecoop/_status CASCADE;`); err != nil {
	// 		return err
	// 	}
	// 	if _, err := db.Exec(`DROP TYPE IF EXISTS github.com/resonatecoop/_group_type CASCADE;`); err != nil {
	// 		return err
	// 	}
	// 	for _, model := range []interface{}{
	// 		// &model.Play{},
	// 		&model.Tag{},
	// 		// &model.github.com/resonatecoop/Group{},
	// 		// &model.github.com/resonatecoop/{},
	// 		&model.GroupTaxonomy{},
	// 		&model.UserGroupMember{},
	// 		&model.StreetAddress{},
	// 		&model.UserGroupPrivacy{},
	// 		&model.UserGroup{},
	// 		&model.User{},
	// 		&model.Link{},
	// 	} {
	// 		err := orm.DropTable(db.(orm.DB), model, &orm.DropTableOptions{
	// 			IfExists: true,
	// 			Cascade:  true,
	// 		})
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// 	return nil
	// })
}

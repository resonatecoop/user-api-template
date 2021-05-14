package main

import (
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"

	"github.com/resonatecoop/user-api/internal/model"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
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

		for _, model := range []interface{}{
			&model.StreetAddress{},
			&model.Tag{},
			&model.User{},
			&model.Link{},
			&model.UserGroupPrivacy{},
			&model.GroupTaxonomy{},
			&model.UserGroup{},
			// &model.github.com/resonatecoop/{},
			// &model.github.com/resonatecoop/Group{},
			&model.UserGroupMember{},
			// &model.Play{},
		} {
			err := orm.CreateTable(db.(orm.DB), model, &orm.CreateTableOptions{
				FKConstraints: true,
				IfNotExists:   true,
			})
			if err != nil {
				return err
			}
		}
		orm.RegisterTable((*model.UserGroupMember)(nil))
		if _, err := db.Exec(`alter table user_group_members add foreign key (user_group_id) references user_groups(id)`); err != nil {
			return err
		}
		if _, err := db.Exec(`alter table user_group_members add foreign key (member_id) references user_groups(id)`); err != nil {
			return err
		}
		return nil
	}, func(db migrations.DB) error {
		if _, err := db.Exec(`DROP TYPE IF EXISTS play_type CASCADE;`); err != nil {
			return err
		}
		if _, err := db.Exec(`DROP TYPE IF EXISTS github.com/resonatecoop/_status CASCADE;`); err != nil {
			return err
		}
		if _, err := db.Exec(`DROP TYPE IF EXISTS github.com/resonatecoop/_group_type CASCADE;`); err != nil {
			return err
		}
		for _, model := range []interface{}{
			// &model.Play{},
			&model.Tag{},
			// &model.github.com/resonatecoop/Group{},
			// &model.github.com/resonatecoop/{},
			&model.GroupTaxonomy{},
			&model.UserGroupMember{},
			&model.StreetAddress{},
			&model.UserGroupPrivacy{},
			&model.UserGroup{},
			&model.User{},
			&model.Link{},
		} {
			err := orm.DropTable(db.(orm.DB), model, &orm.DropTableOptions{
				IfExists: true,
				Cascade:  true,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}

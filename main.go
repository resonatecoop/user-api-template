package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/migrate"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpclog "google.golang.org/grpc/grpclog"

	"github.com/resonatecoop/user-api-template/app"
	"github.com/resonatecoop/user-api-template/migrations"
	"github.com/resonatecoop/user-api-template/model"
	acc "github.com/resonatecoop/user-api-template/pkg/access"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/resonatecoop/user-api-template/gateway"
	"github.com/resonatecoop/user-api-template/insecure"
	pbUser "github.com/resonatecoop/user-api-template/proto/user"

	authorization "github.com/resonatecoop/user-api-template/authorization"
	userserver "github.com/resonatecoop/user-api-template/server"

	// Static files
	_ "github.com/resonatecoop/user-api-template/statik"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "userapi",

		Commands: []*cli.Command{
			runServerCommand,
			newDBCommand(migrations.Migrations),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var runServerCommand = &cli.Command{
	Name:  "runserver",
	Usage: "start User API server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "env",
			Value: "dev",
			Usage: "runtime environment (prod uses env variables for connections)",
		},
		&cli.StringFlag{
			Name:  "dbdebug",
			Value: "false",
			Usage: "show database queries true/false",
		},
	},

	Action: func(c *cli.Context) error {
		_, apiapp, err := app.Start(c.Context, "api", c.String("env"))

		//Adds gRPC internal logs. This is quite verbose, so adjust as desired!
		log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
		grpclog.SetLoggerV2(log)

		checkErr(log, err)

		cfg := apiapp.Cfg

		dbdebug := false

		if c.String("dbdebug") == "true" {
			dbdebug = true
		}

		db := apiapp.DB(c.String("env"), dbdebug)

		accService := acc.New(cfg.Access.NoTokenMethods, cfg.Access.PublicMethods, cfg.Access.WriteMethods)

		interceptorAuth := authorization.NewAuthInterceptor(db, cfg.RefreshToken.Lifetime, accService)

		addr := "0.0.0.0:10000"
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}

		opts := make([]grpc.ServerOption, 0)

		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)))

		opts = append(opts, grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptorAuth.Unary(),
			)))

		s := grpc.NewServer(
			opts...,
		)

		pbUser.RegisterResonateUserServer(s, userserver.New(db))

		// Serve gRPC Server
		log.Info("Serving gRPC on https://", addr)
		go func() {
			log.Fatal(s.Serve(lis))
		}()

		err = gateway.Run("dns:///" + addr)
		log.Fatal(err)

		return err
	},
}

func newDBCommand(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "runtime environment (dev, test, prod) (defaults to dev, prod uses env variables for connections)",
			},
			&cli.StringFlag{
				Name:  "dbdebug",
				Value: "false",
				Usage: "show database queries true/false",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					return migrator.Init(ctx)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					_, err = migrator.Migrate(ctx)

					return err
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					_, err = migrator.Rollback(ctx)

					return err
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					return migrator.Lock(ctx)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					return migrator.Unlock(ctx)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					_, err = migrator.CreateGoMigration(ctx, c.Args().Get(0))

					return err
				},
			},
			{
				Name:  "create_sql",
				Usage: "create SQL migration",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}
					migrator := migrate.NewMigrator(app.DB(c.String("env"), dbdebug), migrations)

					_, err = migrator.CreateSQLMigrations(ctx, c.Args().Get(0))

					return err
				},
			},
			{
				Name:  "load_default_fixtures",
				Usage: "load default data",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					db := app.DB(c.String("env"), dbdebug)

					// Let the db know about the models.
					models := []interface{}{
						(*model.Role)(nil),
						(*model.Scope)(nil),
						(*model.GroupType)(nil),
					}

					for _, this_model := range models {
						db.RegisterModel(this_model)
					}

					fixture := dbfixture.New(db)

					return fixture.Load(ctx, os.DirFS("fixtures/default"), "default_fixtures.yaml")
				},
			},
			{
				Name:  "load_test_fixtures",
				Usage: "load test data",
				Action: func(c *cli.Context) error {
					ctx, app, err := app.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					dbdebug := false

					if c.String("dbdebug") == "true" {
						dbdebug = true
					}

					db := app.DB(c.String("env"), dbdebug)

					// Let the db know about the models.
					models := []interface{}{
						(*model.UserGroup)(nil),
						(*model.StreetAddress)(nil),
						(*model.Tag)(nil),
						(*model.Role)(nil),
						(*model.User)(nil),
						(*model.Link)(nil),
						// (*model.UserGroupPrivacy)(nil),
						(*model.GroupType)(nil),
						// (*model.UserGroupMember)(nil),
						(*model.EmailToken)(nil),
						(*model.Client)(nil),
						(*model.Scope)(nil),
						(*model.RefreshToken)(nil),
						(*model.AuthorizationCode)(nil),
						(*model.AccessToken)(nil),
					}

					for _, this_model := range models {
						db.RegisterModel(this_model)
					}

					fixture := dbfixture.New(db)

					return fixture.Load(ctx, os.DirFS("fixtures/test"), "test_fixtures.yaml")
				},
			},
		},
	}
}

func checkErr(log grpclog.LoggerV2, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

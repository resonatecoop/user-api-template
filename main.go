package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	bun "github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	bundebug "github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpclog "google.golang.org/grpc/grpclog"

	"github.com/resonatecoop/user-api/app"
	"github.com/resonatecoop/user-api/migrations"
	"github.com/resonatecoop/user-api/model"
	acc "github.com/resonatecoop/user-api/pkg/access"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/resonatecoop/user-api/gateway"
	"github.com/resonatecoop/user-api/insecure"
	jwt "github.com/resonatecoop/user-api/pkg/jwt"
	"github.com/resonatecoop/user-api/pkg/zerolog"
	pbIAM "github.com/resonatecoop/user-api/proto/iam"
	pbUser "github.com/resonatecoop/user-api/proto/user"
	iamserver "github.com/resonatecoop/user-api/server/iam"
	iamdb "github.com/resonatecoop/user-api/server/iam/platform/postgres"
	"github.com/resonatecoop/user-api/server/iam/secure"
	userserver "github.com/resonatecoop/user-api/server/users"

	//userserver_test "github.com/resonatecoop/user-api/server/user/userserver_test"

	// Static files
	_ "github.com/resonatecoop/user-api/statik"

	cli "github.com/urfave/cli/v2"
)

// type OauthTestSuite struct {
// 	suite.Suite
// 	cnf *config.Config

// 	service *oauth.Service
// 	clients []*model.Client
// 	users   []*model.User
// 	router  *mux.Router
// }

func main() {
	app := &cli.App{
		Name: "userapi",
		// Flags: []cli.Flag{
		// 	&cli.StringFlag{
		// 		Name:  "env",
		// 		Value: "dev",
		// 		Usage: "environment",
		// 	},
		//},
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
	// Flags: []cli.Flag{
	// 	&cli.StringFlag{
	// 		Name:  "addr",
	// 		Value: "localhost:8000",
	// 		Usage: "serve address",
	// 	},
	// },

	Action: func(c *cli.Context) error {
		_, apiapp, err := app.Start(c.Context, "api", c.String("env"))

		// cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
		// flag.Parse()

		// cfg, err := config.Load(*cfgPath)

		//Adds gRPC internal logs. This is quite verbose, so adjust as desired!
		log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
		grpclog.SetLoggerV2(log)

		checkErr(log, err)

		cfg := apiapp.Cfg

		sqldb, err := sql.Open("pgx", cfg.DB.Dev.PSN)
		// if err != nil {
		// 	panic(err)
		// }

		checkErr(log, err)

		db := bun.NewDB(sqldb, pgdialect.New())

		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))

		//	db, err := pgsql.New(cfg.DB.Dev.PSN, cfg.DB.Dev.LogQueries, cfg.DB.Dev.TimeoutSeconds)

		zerolog := zerolog.New()

		secureSvc := secure.New(cfg.App.MinPasswordStrength)

		jwtService := jwt.New(cfg.JWT.Secret, cfg.JWT.Duration, cfg.JWT.Algorithm)

		accService := acc.New(cfg.Access.NoTokenMethods, cfg.Access.PublicUserMethods)

		iamServer := iamserver.New(db, jwtService, iamdb.NewUser(), secureSvc)

		interceptorAuth := iamserver.NewAuthInterceptor(jwtService, accService)

		addr := "0.0.0.0:10000"
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}

		opts := make([]grpc.ServerOption, 0)

		// TODO: Replace with your own certificate!
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)))

		opts = append(opts, grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptorAuth.Unary(),
			)))

		s := grpc.NewServer(
			opts...,
		)

		pbUser.RegisterResonateUserServer(s, userserver.New(db, secureSvc))

		iamserver.NewLoggingService(iamServer, zerolog)

		pbIAM.RegisterResonateIAMServer(s, iamServer)

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
		Subcommands: []*cli.Command{
			{
				Name:  "dev",
				Usage: "address dev database",
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

							return migrations.Init(ctx, app.DB("dev"))
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

							return migrations.Migrate(ctx, app.DB("dev"))
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

							return migrations.Rollback(ctx, app.DB("dev"))
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

							return migrations.Lock(ctx, app.DB("dev"))
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

							return migrations.Unlock(ctx, app.DB("dev"))
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

							return migrations.CreateGo(ctx, app.DB("dev"), c.Args().Get(0))
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

							return migrations.CreateSQL(ctx, app.DB("dev"), c.Args().Get(0))
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

							cfg := app.Cfg

							sqldb, err := sql.Open("pgx", cfg.DB.Dev.PSN)
							if err != nil {
								return err
							}

							db := bun.NewDB(sqldb, pgdialect.New())

							// Let the db know about the models.
							models := []interface{}{
								(*model.Role)(nil),
								(*model.Scope)(nil),
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

							cfg := app.Cfg

							sqldb, err := sql.Open("pgx", cfg.DB.Dev.PSN)
							if err != nil {
								return err
							}

							db := bun.NewDB(sqldb, pgdialect.New())

							// Let the db know about the models.
							models := []interface{}{
								(*model.UserGroup)(nil),
								(*model.StreetAddress)(nil),
								(*model.Tag)(nil),
								(*model.Role)(nil),
								(*model.User)(nil),
								(*model.Link)(nil),
								(*model.UserGroupPrivacy)(nil),
								(*model.GroupTaxonomy)(nil),
								(*model.UserGroupMember)(nil),
								(*model.EmailToken)(nil),
								(*model.EmailTokenClaims)(nil),
								(*model.Email)(nil),
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
			},
			{
				Name:  "test",
				Usage: "address test database",
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

							return migrations.Init(ctx, app.DB("test"))
						},
					},
					{
						Name:  "migrate",
						Usage: "migrate test database",
						Action: func(c *cli.Context) error {
							ctx, app, err := app.StartCLI(c)
							if err != nil {
								return err
							}
							defer app.Stop()

							return migrations.Migrate(ctx, app.DB("test"))
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback the last migration group on test db",
						Action: func(c *cli.Context) error {
							ctx, app, err := app.StartCLI(c)
							if err != nil {
								return err
							}
							defer app.Stop()

							return migrations.Rollback(ctx, app.DB("test"))
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

							return migrations.Lock(ctx, app.DB("test"))
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

							return migrations.Unlock(ctx, app.DB("test"))
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

							return migrations.CreateGo(ctx, app.DB("test"), c.Args().Get(0))
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

							return migrations.CreateSQL(ctx, app.DB("test"), c.Args().Get(0))
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

							cfg := app.Cfg

							sqldb, err := sql.Open("pgx", cfg.DB.Test.PSN)
							if err != nil {
								return err
							}

							db := bun.NewDB(sqldb, pgdialect.New())

							// Let the db know about the models.
							models := []interface{}{
								(*model.Role)(nil),
								(*model.Scope)(nil),
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

							cfg := app.Cfg

							sqldb, err := sql.Open("pgx", cfg.DB.Test.PSN)
							if err != nil {
								return err
							}

							db := bun.NewDB(sqldb, pgdialect.New())

							// Let the db know about the models.
							models := []interface{}{
								(*model.UserGroup)(nil),
								(*model.StreetAddress)(nil),
								(*model.Tag)(nil),
								(*model.Role)(nil),
								(*model.User)(nil),
								(*model.Link)(nil),
								(*model.UserGroupPrivacy)(nil),
								(*model.GroupTaxonomy)(nil),
								(*model.UserGroupMember)(nil),
								(*model.EmailToken)(nil),
								(*model.EmailTokenClaims)(nil),
								(*model.Email)(nil),
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
			},
		},
	}
}

func checkErr(log grpclog.LoggerV2, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

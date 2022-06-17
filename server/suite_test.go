package server_test

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"testing"

	//"github.com/resonatecoop/id/log"
	// "github.com/resonatecoop/id/log"
	"github.com/resonatecoop/user-api-template/model"
	"github.com/resonatecoop/user-api-template/pkg/config"
	"github.com/resonatecoop/user-api-template/server"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
// testDbUser = "go_oauth2_server"
// testDbName = "go_oauth2_server_oauth_test"

// testFixtures = []string{
// 	"./oauth/fixtures/scopes.yml",
// 	"./oauth/fixtures/roles.yml",
// 	"./oauth/fixtures/test_clients.yml",
// 	"./oauth/fixtures/test_users.yml",
// }

// 	testMigrations = []func(*bun.DB) error{
// 		model.MigrateAll,
// 	}
)

func init() {
	if err := os.Chdir("../"); err != nil {
		log.Fatal(err)
	}
}

// UserApiTestSuite needs to be exported so the tests run
type UserApiTestSuite struct {
	suite.Suite
	cfg    *config.Configuration
	db     *bun.DB
	ctx    context.Context
	server *server.Server
	// service *oauth.Service
	// clients []*model.Client
	// users   []*model.User
	// router  *mux.Router
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *UserApiTestSuite) SetupSuite() {
	// Initialise the config
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)

	if err != nil {
		panic(err)
	}

	sqldb, err := sql.Open("pgx", cfg.DB.Dev.PSN)

	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	suite.cfg = cfg
	suite.ctx = context.Background()
	suite.db = db

	suite.server = server.New(db)

	// if err != nil {
	// 	panic(err)
	// }

	// ASSUME THAT TEST DATABASE HAS ALREADY BEEN CREATED
	// Create the test database
	// db, err := testutil.CreateTestDatabasePostgres(
	// 	suite.cnf.Database.Host,
	// 	testDbUser,
	// 	testDbName,
	// 	testMigrations,
	// 	testFixtures,
	// )
	// if err != nil {
	// 	t.Skip(err)
	// }
	// suite.db = db
	// suite.db2 = nil // TODO setup test mysql db client

	// Fetch test client
	// suite.clients = make([]*model.Client, 0)

	// rows, err := suite.db.QueryContext(ctx, "SELECT * from clients ORDER BY created_at")

	// if err != nil {
	// 	panic(err)
	// }

	// err = suite.db.ScanRows(ctx, rows, &suite.clients)

	// if err != nil {
	// 	log.ERROR.Fatal(err)
	// }

	// // Fetch test users
	// suite.users = make([]*model.User, 0)

	// rows, err = suite.db.QueryContext(ctx, "SELECT * from users ORDER BY created_at")

	// if err != nil {
	// 	panic(err)
	// }

	// err = suite.db.ScanRows(ctx, rows, &suite.users)

	// if err != nil {
	// 	log.ERROR.Fatal(err)
	// }
	// // Initialise the service
	// suite.service = oauth.NewService(suite.cnf, suite.db)

	// // Register routes
	// suite.router = mux.NewRouter()
	// suite.service.RegisterRoutes(suite.router, "/v1/oauth")

	// return nil
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (suite *UserApiTestSuite) TearDownSuite() {
	//
}

// The SetupTest method will be run before every test in the suite.
func (suite *UserApiTestSuite) SetupTest() {
	//
}

// The TearDownTest method will be run after every test in the suite.
func (suite *UserApiTestSuite) TearDownTest() {
	// Scopes are static, populated from fixtures,
	// so there is no need to clear them after running a test

	suite.db.NewTruncateTable().
		Model(new(model.AuthorizationCode)).
		Exec(suite.ctx)

	suite.db.NewTruncateTable().
		Model(new(model.RefreshToken)).
		Exec(suite.ctx)

	suite.db.NewTruncateTable().
		Model(new(model.AccessToken)).
		Exec(suite.ctx)

	ids := []string{
		"243b4178-6f98-4bf1-bbb1-46b57a901816",
		"5253747c-2b8c-40e2-8a70-bab91348a9bd",
		"90b26113-37e0-456a-9f75-01db0eb550f8",
		"f40cf437-eef2-4659-8eb3-7ee93f6dfcea",
		"046e7a23-a9f0-4fb8-a60f-861b2df05d95",
		"953c3537-3149-4164-9c73-ec785a3c17c4"}

	suite.db.NewDelete().
		Model(new(model.User)).
		Where("id NOT IN (?)", bun.In(ids)).
		ForceDelete().
		Exec(suite.ctx)

	suite.db.NewUpdate().
		Model(new(model.User)).
		Set("deleted_at = NULL").
		WhereAllWithDeleted().
		Where("deleted_at IS NOT NULL").
		Exec(suite.ctx)

	ids = []string{"3392e754-ba3e-424f-a687-add9a8ab39c9", "295be195-898c-4f0c-b6a0-8c62105f42de"}

	suite.db.NewDelete().
		Model(new(model.Client)).
		Where("id NOT IN (?)", bun.In(ids)).
		ForceDelete().
		Exec(suite.ctx)
}

// TestUserApiTestSuite ...
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUserApiTestSuite(t *testing.T) {
	suite.Run(t, new(UserApiTestSuite))
}

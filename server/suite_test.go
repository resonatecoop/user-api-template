package oauth_test

import (
	"context"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/resonatecoop/id/config"
	"github.com/resonatecoop/id/database"
	"github.com/resonatecoop/id/log"
	"github.com/resonatecoop/id/oauth"
	"github.com/resonatecoop/user-api/model"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
)

var (
	testDbUser = "go_oauth2_server"
	testDbName = "go_oauth2_server_oauth_test"

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
		log.ERROR.Fatal(err)
	}
}

// OauthTestSuite needs to be exported so the tests run
type OauthTestSuite struct {
	suite.Suite
	cnf     *config.Config
	db      *bun.DB
	service *oauth.Service
	clients []*model.Client
	users   []*model.User
	router  *mux.Router
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *OauthTestSuite) SetupSuite() {
	// Initialise the config
	suite.cnf = config.NewConfig(false, false, "etcd")

	var err error

	suite.db, err = database.NewDatabase(suite.cnf)

	if err != nil {
		panic(err)
	}

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

	ctx := context.Background()
	// Fetch test client
	suite.clients = make([]*model.Client, 0)

	rows, err := suite.db.QueryContext(ctx, "SELECT * from clients ORDER BY created_at")

	if err != nil {
		panic(err)
	}

	err = suite.db.ScanRows(ctx, rows, &suite.clients)

	if err != nil {
		log.ERROR.Fatal(err)
	}

	// Fetch test users
	suite.users = make([]*model.User, 0)

	rows, err = suite.db.QueryContext(ctx, "SELECT * from users ORDER BY created_at")

	if err != nil {
		panic(err)
	}

	err = suite.db.ScanRows(ctx, rows, &suite.users)

	if err != nil {
		log.ERROR.Fatal(err)
	}
	// Initialise the service
	suite.service = oauth.NewService(suite.cnf, suite.db)

	// Register routes
	suite.router = mux.NewRouter()
	suite.service.RegisterRoutes(suite.router, "/v1/oauth")
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (suite *OauthTestSuite) TearDownSuite() {
	//
}

// The SetupTest method will be run before every test in the suite.
func (suite *OauthTestSuite) SetupTest() {
	//
}

// The TearDownTest method will be run after every test in the suite.
func (suite *OauthTestSuite) TearDownTest() {
	// Scopes are static, populated from fixtures,
	// so there is no need to clear them after running a test
	ctx := context.Background()

	suite.db.NewTruncateTable().
		Model(new(model.AuthorizationCode)).
		Exec(ctx)

	suite.db.NewTruncateTable().
		Model(new(model.RefreshToken)).
		Exec(ctx)

	suite.db.NewTruncateTable().
		Model(new(model.AccessToken)).
		Exec(ctx)

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
		Exec(ctx)

	ids = []string{"3392e754-ba3e-424f-a687-add9a8ab39c9", "295be195-898c-4f0c-b6a0-8c62105f42de"}

	suite.db.NewDelete().
		Model(new(model.Client)).
		Where("id NOT IN (?)", bun.In(ids)).
		Exec(ctx)
}

// TestOauthTestSuite ...
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOauthTestSuite(t *testing.T) {
	suite.Run(t, new(OauthTestSuite))
}

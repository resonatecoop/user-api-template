package app

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4"
	stdlib "github.com/jackc/pgx/v4/stdlib"
	bun "github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/resonatecoop/user-api-template/pkg/config"

	//"github.com/uptrace/bun/extra/bundebug"
	"github.com/urfave/cli/v2"
	//	"github.com/vmihailenco/treemux"
	grpclog "google.golang.org/grpc/grpclog"
)

type appCtxKey struct{}

func AppFromContext(ctx context.Context) *App {
	return ctx.Value(appCtxKey{}).(*App)
}

func ContextWithApp(ctx context.Context, app *App) context.Context {
	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return ctx
}

type App struct {
	ctx context.Context
	Cfg *config.Configuration

	// log grpclog.NewLoggerV2

	stopping uint32
	stopCh   chan struct{}

	onStop      appHooks
	onAfterStop appHooks

	// router    *treemux.Router
	// apiRouter *treemux.Group

	// lazy init
	dbOnce sync.Once
	db     *bun.DB
}

func New(ctx context.Context, cfg *config.Configuration) *App {

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	// log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	// grpclog.SetLoggerV2(log)

	app := &App{
		// log:    log,
		Cfg:    cfg,
		stopCh: make(chan struct{}),
	}
	app.ctx = ContextWithApp(ctx, app)
	//app.initRouter()
	return app
}

func StartCLI(c *cli.Context) (context.Context, *App, error) {
	return Start(c.Context, c.Command.Name, c.String("env"))
}

func Start(ctx context.Context, service, envName string) (context.Context, *App, error) {

	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)

	if err != nil {
		return nil, nil, err
	}
	return StartConfig(ctx, cfg)
}

func StartConfig(ctx context.Context, cfg *config.Configuration) (context.Context, *App, error) {
	rand.Seed(time.Now().UnixNano())

	app := New(ctx, cfg)
	if err := onStart.Run(ctx, app); err != nil {
		return nil, nil, err
	}
	return app.ctx, app, nil
}

func (app *App) Stop() {
	_ = app.onStop.Run(app.ctx, app)
	_ = app.onAfterStop.Run(app.ctx, app)
}

func (app *App) OnStop(name string, fn HookFunc) {
	app.onStop.Add(newHook(name, fn))
}

func (app *App) OnAfterStop(name string, fn HookFunc) {
	app.onAfterStop.Add(newHook(name, fn))
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Config() *config.Configuration {
	return app.Cfg
}

func (app *App) Running() bool {
	return !app.Stopping()
}

func (app *App) Stopping() bool {
	return atomic.LoadUint32(&app.stopping) == 1
}

// func (app *App) IsDebug() bool {
// 	return app.cfg.Debug
// }

// func (app *App) Router() *treemux.Router {
// 	return app.router
// }

// func (app *App) APIRouter() *treemux.Group {
// 	return app.apiRouter
// }

func (app *App) DB(env string, debug bool) *bun.DB {
	app.dbOnce.Do(func() {

		var err error

		var db *bun.DB

		var dbPSN string

		databaseName := os.Getenv("POSTGRES_NAME")

		databaseUser := os.Getenv("POSTGRES_USER")

		if databaseName != "" && databaseUser != "" {
			databaseHost := os.Getenv("POSTGRES_HOST")

			if databaseHost == "" {
				databaseHost = "127.0.0.1"
			}

			databasePort := os.Getenv("POSTGRES_PORT")

			if databasePort == "" {
				databasePort = "5432"
			}

			databaseSslMode := os.Getenv("POSTGRES_SSL")

			if databaseSslMode != "enable" {
				databaseSslMode = "disable"
			}

			databasePass := os.Getenv("POSTGRES_PASS")

			dbPSN = "postgres://" + databaseUser + ":" + databasePass + "@" + databaseHost + ":" + databasePort + "/" + databaseName + "?sslmode=" + databaseSslMode
		}

		switch env {
		case "test":
			dbPSN = app.Cfg.DB.Test.PSN
		case "dev":
			dbPSN = app.Cfg.DB.Dev.PSN
		case "prod":
			//do nothing, should be set up above
			break
		default:
			dbPSN = app.Cfg.DB.Dev.PSN
		}

		db = connectDB(dbPSN, debug)

		if err != nil {
			panic(err)
		}

		app.OnStop("db.Close", func(ctx context.Context, _ *App) error {
			return db.Close()
		})

		app.db = db

	})
	return app.db
}

//------------------------------------------------------------------------------

func WaitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}

func checkErr(log grpclog.LoggerV2, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectDB(PSN string, isDebug bool) *bun.DB {

	dbconfig, err := pgx.ParseConfig(PSN)

	if err != nil {
		panic(err)
	}

	dbconfig.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*dbconfig)

	db := bun.NewDB(sqldb, pgdialect.New())
	if isDebug {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return db
}

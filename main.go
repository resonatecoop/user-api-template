package main

import (
	"flag"
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpclog "google.golang.org/grpc/grpclog"

	"github.com/merefield/grpc-user-api/gateway"
	"github.com/merefield/grpc-user-api/insecure"
	"github.com/merefield/grpc-user-api/pkg/config"
	"github.com/merefield/grpc-user-api/pkg/jwt"
	pgsql "github.com/merefield/grpc-user-api/pkg/postgres"
	"github.com/merefield/grpc-user-api/pkg/zerolog"
	pbIAM "github.com/merefield/grpc-user-api/proto/iam"
	pbUser "github.com/merefield/grpc-user-api/proto/user"
	iamserver "github.com/merefield/grpc-user-api/server/iam"
	iamdb "github.com/merefield/grpc-user-api/server/iam/platform/postgres"
	"github.com/merefield/grpc-user-api/server/iam/secure"
	userserver "github.com/merefield/grpc-user-api/server/user"

	// Static files
	_ "github.com/merefield/grpc-user-api/statik"
)

func main() {

	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	checkErr(log, err)

	db, err := pgsql.New(cfg.DB.Dev.PSN, cfg.DB.Dev.LogQueries, cfg.DB.Dev.TimeoutSeconds)

	checkErr(log, err)

	zerolog := zerolog.New()

	secureSvc := secure.New(cfg.App.MinPasswordStrength)

	j := jwt.New(cfg.JWT.Secret, cfg.JWT.Duration, cfg.JWT.Algorithm)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	pbUser.RegisterResonateUserServer(s, userserver.New(db, secureSvc))

	iams := iamserver.New(db, j, iamdb.NewUser(), secureSvc)

	iamserver.NewLoggingService(iams, zerolog)

	pbIAM.RegisterResonateIAMServer(s, iams)
	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatal(err)
}

func checkErr(log grpclog.LoggerV2, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	pgsql "github.com/merefield/grpc-user-api/pkg/postgres"
	pbAPI "github.com/merefield/grpc-user-api/proto"
	"github.com/merefield/grpc-user-api/server"

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

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	pbAPI.RegisterUserServiceServer(s, server.New(db))

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

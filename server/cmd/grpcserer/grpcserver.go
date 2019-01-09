package main

import (
	"../../internal/_env"
	"../../internal/_healthcheck"
	"../../internal/auth"
	"../../internal/task"
	"./middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

const (
	port = ":9090"
)

func main() {
	dbSession, err := configureDB()
	defer dbSession.Close()

	// config kann hier nachher an env angehängt werden
	env := &environment.Env{DB: dbSession}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_auth.UnaryServerInterceptor(middleware.JWTAuthFunc))))

	// Auth
	auth.InitEnvironment(env)
	auth.RegisterAuthServiceServer(grpcServer, auth.GetServiceServer())

	// Healthcheck
	_healthcheck.RegisterHealthCheckServer(grpcServer, _healthcheck.GetServiceServer())

	// Task
	task.InitEnvironment(env)
	task.RegisterTaskServiceServer(grpcServer, task.GetServiceServer())

	// weitere Services kann man hier registrieren

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func configureDB() (sqlbuilder.Database, error) {
	var settings = sqlite.ConnectionURL{
		Database: `data/data.db`, // Path to database file.
	}
	dbSession, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	return dbSession, err
}

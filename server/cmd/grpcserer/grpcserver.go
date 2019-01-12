package main

import (
	"../../internal/_healthcheck"
	"../../internal/auth"
	"../../internal/pkg/environment"
	"../../internal/tag"
	"../../internal/task"
	"./middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":9090"
)

func main() {

	environment.InitEnv()
	defer environment.Env.DB.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_auth.UnaryServerInterceptor(middleware.JWTAuthFunc))))

	// Auth
	auth.Register()
	auth.RegisterAuthServiceServer(grpcServer, auth.GetServiceServer())

	// Healthcheck
	//_healthcheck.Register()
	_healthcheck.RegisterHealthCheckServer(grpcServer, _healthcheck.GetServiceServer())

	// Task
	task.Register()
	task.RegisterServiceServer(grpcServer, task.GetServiceServer())

	// Tag
	tag.Register()
	tag.RegisterServiceServer(grpcServer, tag.GetServiceServer())

	// weitere Services kann man hier registrieren

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

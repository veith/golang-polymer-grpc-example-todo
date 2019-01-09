package _healthcheck

import (
	proto "../../../proto/healthcheck"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
)

var RegisterHealthCheckServer = proto.RegisterHealthCheckServer

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.HealthCheckServer {
	var s healthcheckServiceServer
	return &s
}

// healthcheckServiceServer is used to implement healthcheckServiceServer.
type healthcheckServiceServer struct {
}

func (healthcheckServiceServer) Check(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	// Todo: verbindung auf db ev. auch prüfen und gegebenfalls error zurück geben
	fmt.Println(ctx.Value("tokenInfo"))
	return &empty.Empty{}, nil
}

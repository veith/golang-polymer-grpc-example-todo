package main

import (
	"../../../proto/auth"
	"../../../proto/healthcheck"
	"../../../proto/task"
	"../../../proto/tag"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

var (
	//Todo: Port und Adresse des Transcoders via Flags setzbar machen.
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of YourService")
)

// header für client ohne prefixes
func outgoingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

// header vom client im ctx des Servers ohne prefixes senden
func incomingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithOutgoingHeaderMatcher(outgoingMatcher), runtime.WithIncomingHeaderMatcher(incomingMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Services hier einfügen
	// Auth Service
	err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	//Healthcheck
	err = healthcheck.RegisterHealthCheckHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	//Tasks
	err = task.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	//Tag
	err = tag.RegisterTagServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8181", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

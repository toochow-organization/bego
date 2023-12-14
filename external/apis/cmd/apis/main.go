package main

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/toochow-organization/bego/base"
	"github.com/toochow-organization/bego/base/config"
	"github.com/toochow-organization/bego/base/id"
	"github.com/toochow-organization/bego/base/log"
	"google.golang.org/grpc"

	apis "github.com/toochow-organization/bego/external/apis"
	"github.com/toochow-organization/bego/external/apis/inmem"
	pb "github.com/toochow-organization/bego/protocol/external/apis/v1"
)

func main() {
	// Create a random service name if not provided
	serviceName := config.LookupEnv("BEGO_SERVICE_NAME", id.NewGenerator("apis").Generate())
	// Init context
	ctx := context.Background()
	// Initiate a logger with pre-configuration for production and telemetry.
	l, err := log.New()
	if err != nil {
		// in case we cannot create the logger, the app should immediately stop.
		panic(err)
	}
	// Replace the global logger with the Service scoped log.
	log.ReplaceGlobal(l)

	// Initialise the boilerplate and start the service
	boilerplate, err := base.NewBoilerPlate(serviceName, base.WithLogger(l))
	if err != nil {
		l.Fatal(ctx, err.Error())
	}
	// Setup for version service (sample)
	storage := inmem.NewVersionStorage()
	svc := apis.NewVersionService(storage)
	srv := newHandler(svc)

	// Register the GRPC Server
	boilerplate.RegisterService(func(s *grpc.Server) {
		pb.RegisterApiServiceServer(s, srv)
	})

	// Register the Service Handler
	boilerplate.RegisterServiceHandler(func(gw *runtime.ServeMux, conn *grpc.ClientConn) {
		if err := pb.RegisterApiServiceHandler(ctx, gw, conn); err != nil {
			l.Fatal(ctx, "fail registering gateway handler", log.Error(err))
		}
	})

	l.Info(ctx, "Starting service", log.String("service.name", serviceName))
	if err := boilerplate.Start(); err != nil {
		l.Error(ctx, "fail starting", log.Error(err))
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/toochow-organization/bego/base"
	"github.com/toochow-organization/bego/base/config"
	"github.com/toochow-organization/bego/base/id"
	"github.com/toochow-organization/bego/base/log"
	appPg "github.com/toochow-organization/bego/external/apis/postgres"
	"google.golang.org/grpc"

	// gorm postgres
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	apis "github.com/toochow-organization/bego/external/apis"
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

	// configuration
	conf := configuration{}
	err = config.FromConfigMap(&conf)
	if err != nil {
		l.Fatal(ctx, err.Error())
	}

	// Initialise the boilerplate and start the service
	core, err := base.NewCore(serviceName, base.WithLogger(l))
	if err != nil {
		l.Fatal(ctx, err.Error())
	}

	// Init DB, for now, we use gorm
	// @TODO: make a wrapper for gorm in the base folder
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		conf.Username,
		conf.Password,
		conf.HostAndPort,
		conf.Database,
		conf.SSLMode)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		l.Fatal(ctx, err.Error())
	}
	// Setup for version service (sample)
	modelStorage := appPg.NewModelStorage(l, db, appPg.ShouldMigrateModelStorage(conf.Migrate))
	svc := apis.NewModelService(modelStorage)
	srv := newHandler(l, svc)

	// Register the GRPC Server
	core.RegisterService(func(s *grpc.Server) {
		pb.RegisterApiServiceServer(s, srv)
	})

	// Register the Service Handler
	core.RegisterServiceHandler(func(gw *runtime.ServeMux, conn *grpc.ClientConn) {
		if err := pb.RegisterApiServiceHandler(ctx, gw, conn); err != nil {
			l.Fatal(ctx, "fail registering gateway handler", log.Error(err))
		}
	})

	l.Info(ctx, "Starting service", log.String("service.name", serviceName))
	if err := core.Start(); err != nil {
		l.Error(ctx, "fail starting", log.Error(err))
	}
}

package base

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	// chi as router
	mux "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/toochow-organization/bego/base/config"
	"github.com/toochow-organization/bego/base/errors"
	log "github.com/toochow-organization/bego/base/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// Core provides a set of common base code for creating a production ready GRPC server and
// HTTP mux router (with grpc-gateway capabilities) and custom HTTP endpoints
type Core struct {
	// Base component's name
	name string
	// Options
	opts *CoreOptions
	// logger
	logger *log.Logger
	// Runtime (grpc-gateway)
	runtimeServer   *runtime.ServeMux
	runtimeClient   *grpc.ClientConn
	runtimeInstance sync.Once
	// gRPC server
	grpcServer   *grpc.Server
	grpcInstance sync.Once
	// HTTP server
	liveness     http.HandlerFunc
	readiness    http.HandlerFunc
	httpServer   *http.Server
	httpHandler  *mux.Mux
	httpInstance sync.Once
}

var defaultHealthHandler = func(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(writer, "ok") //nolint
}

func NewCore(name string, opts ...func(*CoreOptions)) (*Core, error) {
	if len(name) == 0 {
		return nil, errors.New("name cannot be empty")
	}
	CoreOpts := &CoreOptions{
		httpAddr: config.LookupEnv("BEGO_HTTP_ADDR", "0.0.0.0:8080"),
		grpcAddr: config.LookupEnv("BEGO_GRPC_ADDR", "0.0.0.0:8081"),
		// By default, always 15 seconds timeout
		httpWriteTimeout: 15 * time.Second,
		httpReadTimeout:  15 * time.Second,
		logger:           log.NewNop(),
	}
	for _, o := range opts {
		o(CoreOpts)
	}

	return &Core{
		name:      name,
		opts:      CoreOpts,
		logger:    CoreOpts.logger,
		readiness: defaultHealthHandler,
		liveness:  defaultHealthHandler,
	}, nil
}

func (b *Core) initHTTPOnce() {
	// make sure the HTTP server has been initialized
	b.httpInstance.Do(func() {
		router := mux.NewRouter()
		// Gzip compression for clients that accept compressed responses
		// Passing a compression level of 5 is sensible value
		router.Use(middleware.Compress(5))
		// @TODO: add tracer and metrics middleware when needed here

		// If cors is enabled, we should set it depending on the options
		if b.opts.enableCors {
			router.Use(cors.New(b.opts.corsOpts).Handler)
		}
		b.httpHandler = router
		b.httpServer = &http.Server{
			Addr:         b.opts.httpAddr,
			Handler:      b.httpHandler,
			WriteTimeout: b.opts.httpWriteTimeout,
			ReadTimeout:  b.opts.httpReadTimeout,
		}
	})
}

// RegisterServiceFunc represents a function for registering a grpc service handler.
type RegisterServiceFunc func(s *grpc.Server)

// RegisterService registers a grpc service handler.
func (b *Core) RegisterService(fn RegisterServiceFunc) {
	// Create GRPC server only once
	b.grpcInstance.Do(func() {
		b.grpcServer = b.newGrpcSever(b.opts.grpcServerOpts...)
	})
	fn(b.grpcServer)
}

// newGrpcSever creates a new gRPC server with the provided options.
// It applies stream and unary interceptors for recovery and validation.
// Additional server options can be passed as variadic arguments.
// Returns the created gRPC server.
func (b *Core) newGrpcSever(opts ...grpc.ServerOption) *grpc.Server {
	serverOpts := []grpc.ServerOption{
		grpc.ChainStreamInterceptor(
			grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandlerContext(recoverFrom(log.L()))),
			grpcValidator.StreamServerInterceptor(),
		),
		grpc.ChainUnaryInterceptor(
			grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandlerContext(recoverFrom(log.L()))),
			grpcValidator.UnaryServerInterceptor(),
		),
	}

	serverOpts = append(serverOpts, opts...)
	return grpc.NewServer(serverOpts...)
}

func recoverFrom(l *log.Logger) grpcRecovery.RecoveryHandlerFuncContext {
	return func(ctx context.Context, p interface{}) error {
		l.Error(ctx, "grpc recover panic", zap.Any("panic", p))
		return status.Errorf(codes.Internal, "%v", p)
	}
}

func (b *Core) newGrprClient() (*grpc.ClientConn, error) {
	dialOps := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			grpcRetry.UnaryClientInterceptor(),
			grpcValidator.UnaryClientInterceptor(),
		),
		grpc.WithChainStreamInterceptor(
			grpcRetry.StreamClientInterceptor(),
		),
	}
	// We can add more grpc dial options here
	// such as telemetry, tracing, etc.
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	dialOps = append(dialOps, grpcOpts...)
	return grpc.Dial(b.opts.grpcAddr, dialOps...)
}

func (b *Core) initRuntimeOnce(muxOpts ...runtime.ServeMuxOption) {
	b.runtimeInstance.Do(func() {
		b.logger.Info(context.Background(), "initializing grpc-gateway")

		conn, err := b.newGrprClient()
		if err != nil {
			b.logger.Fatal(context.Background(), "failed to dial server", log.Error(err))
		}
		b.runtimeClient = conn
		// multiplexer options
		// we can add more handler options here
		muxOpts = append(
			muxOpts,
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true,
						EmitUnpopulated: true,
					},
				},
			}),
		)
		b.runtimeServer = runtime.NewServeMux(muxOpts...)
	})
}

// RegisterServiceHandlerFunc represents a function for registering a grpc gateway service handler.
type RegisterServiceHandlerFunc func(gw *runtime.ServeMux, conn *grpc.ClientConn)

// RegisterServiceHandler registers a grpc-gateway service handler.
// Reference: https://github.com/grpc-ecosystem/grpc-gateway
func (b *Core) RegisterServiceHandler(fn RegisterServiceHandlerFunc, muxOpts ...runtime.ServeMuxOption) {
	b.initHTTPOnce()
	// Only create one time the gateway and grpc client
	b.initRuntimeOnce(muxOpts...)
	fn(b.runtimeServer, b.runtimeClient)
}

// Start serving request via HTTP and GRPC
func (b *Core) Start() error {
	// Package maxprocs lets Go programs easily configure runtime.GOMAXPROCS to match the configured Linux CPU quota.
	// Unlike the top-level automaxprocs package,
	// it lets the caller configure logging and handle errors.
	_, err := maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		b.logger.Info(context.Background(), fmt.Sprintf(s, i))
	}))
	if err != nil {
		return errors.Wrap(err, "setup maxprocs")
	}
	r := mux.NewRouter()
	// Init default health checks.
	r.Method("GET", "/healthz", b.liveness)
	r.Method("GET", "/readyz", b.readiness)
	// shutdown channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverError := make(chan error, 1)
	// start the grpc server
	go func(serverError chan error) {
		// No GRPC server set up.
		if b.grpcServer == nil {
			return
		}
		// Create listener for the grpc server
		listen, err := net.Listen("tcp", b.opts.grpcAddr)
		if err != nil {
			serverError <- errors.Wrap(err, "init net listener")
		}
		serverError <- b.grpcServer.Serve(listen)
		_ = listen.Close() //nolint
	}(serverError)

	// start the http server
	go func(serverError chan error) {
		// No HTTP server set up.
		if b.httpServer == nil {
			return
		}

		// init the http server
		if b.runtimeServer != nil {
			b.httpHandler.Mount("/", b.runtimeServer)
		}
		serverError <- b.httpServer.ListenAndServe()
	}(serverError)

	b.logger.Debug(context.Background(), "service started",
		log.String("service-name", b.name))

	select {
	case err := <-serverError:
		return errors.Wrap(err, "server error")
	case <-shutdown:

		// Terminate GRPC server if started
		if b.grpcServer != nil {
			b.grpcServer.GracefulStop()
		}

		// terminate the HTTP server if started.
		if b.httpServer != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			_ = b.httpServer.Shutdown(ctx) //nolint
		}
	}

	return nil
}

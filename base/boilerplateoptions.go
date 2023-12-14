package base

import (
	"time"

	"github.com/go-chi/cors"
	log "github.com/toochow-organization/bego/base/log"
	"google.golang.org/grpc"
)

type BoilerplateOptions struct {
	httpAddr         string
	grpcAddr         string
	grpcServerOpts   []grpc.ServerOption
	corsOpts         cors.Options
	enableCors       bool
	httpWriteTimeout time.Duration
	httpReadTimeout  time.Duration
	logger           *log.Logger
}

// Option defines a Foundation option.
type Option func(*BoilerplateOptions)

// EnableCors will add cors support to the http server.
func EnableCors() Option {
	return func(bo *BoilerplateOptions) {
		bo.enableCors = true
	}
}

// WithGrpcServerOptions defines GRPC server options.
func WithGrpcServerOptions(opts ...grpc.ServerOption) Option {
	return func(bo *BoilerplateOptions) {
		bo.grpcServerOpts = opts
	}
}

// WithCorsOptions defines http server cors options.
func WithCorsOptions(opts cors.Options) Option {
	return func(bo *BoilerplateOptions) {
		bo.corsOpts = opts
	}
}

// WithHTTPAddr defines a HTTP server host and port.
func WithHTTPAddr(addr string) Option {
	return func(bo *BoilerplateOptions) {
		bo.httpAddr = addr
	}
}

// WithHTTPWriteTimeout defines write timeout for the HTTP server.
func WithHTTPWriteTimeout(timeout time.Duration) Option {
	return func(bo *BoilerplateOptions) {
		bo.httpWriteTimeout = timeout
	}
}

// WithHTTPReadTimeout defines read timeout for the HTTP server.
func WithHTTPReadTimeout(timeout time.Duration) Option {
	return func(bo *BoilerplateOptions) {
		bo.httpReadTimeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(bo *BoilerplateOptions) {
		bo.logger = logger
	}
}

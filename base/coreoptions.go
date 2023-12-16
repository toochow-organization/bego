package base

import (
	"time"

	"github.com/go-chi/cors"
	log "github.com/toochow-organization/bego/base/log"
	"google.golang.org/grpc"
)

type CoreOptions struct {
	httpAddr         string
	grpcAddr         string
	statusAddr       string
	grpcServerOpts   []grpc.ServerOption
	corsOpts         cors.Options
	enableCors       bool
	httpWriteTimeout time.Duration
	httpReadTimeout  time.Duration
	logger           *log.Logger
}

type Option func(*CoreOptions)

func EnableCors() Option {
	return func(co *CoreOptions) {
		co.enableCors = true
	}
}

func WithStatusAddr(addr string) Option {
	return func(co *CoreOptions) {
		co.statusAddr = addr
	}
}

func WithGrpcServerOptions(opts ...grpc.ServerOption) Option {
	return func(co *CoreOptions) {
		co.grpcServerOpts = opts
	}
}

func WithCorsOptions(opts cors.Options) Option {
	return func(co *CoreOptions) {
		co.corsOpts = opts
	}
}

func WithHTTPAddr(addr string) Option {
	return func(co *CoreOptions) {
		co.httpAddr = addr
	}
}

func WithHTTPWriteTimeout(timeout time.Duration) Option {
	return func(co *CoreOptions) {
		co.httpWriteTimeout = timeout
	}
}

func WithHTTPReadTimeout(timeout time.Duration) Option {
	return func(co *CoreOptions) {
		co.httpReadTimeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(co *CoreOptions) {
		co.logger = logger
	}
}

package main

import (
	"context"

	"github.com/toochow-organization/bego/external/apis"
	pb "github.com/toochow-organization/bego/protocol/external/apis/v1"
)

type Server struct {
	pb.UnsafeApiServiceServer
	service apis.VersionService
}

func newHandler(service apis.VersionService) Server {
	return Server{
		service: service,
	}
}

// GetVersion
//
// GetVersion returns the version of the API
func (s Server) GetVersion(ctx context.Context, req *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	return &pb.GetVersionResponse{
		Version: s.service.GetVersion(ctx),
	}, nil
}

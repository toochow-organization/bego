package main

import (
	"context"

	errors "github.com/toochow-organization/bego/base/errors"
	log "github.com/toochow-organization/bego/base/log"
	"github.com/toochow-organization/bego/external/apis"
	"github.com/toochow-organization/bego/protocol/errdetails"
	pb "github.com/toochow-organization/bego/protocol/external/apis/v1"
	"google.golang.org/grpc/codes"
)

type Server struct {
	pb.UnsafeApiServiceServer
	l       *log.Logger
	service apis.ModelService
}

func newHandler(l *log.Logger, service apis.ModelService) Server {
	return Server{
		l:       l,
		service: service,
	}
}

// GetVersion
//
// GetVersion returns the version of the API
func (s Server) ListModels(ctx context.Context, req *pb.ListModelsRequest) (*pb.ListModelsResponse, error) {
	models, err := s.service.ListModels(ctx)
	if err != nil {
		// log error
		s.l.Error(ctx, err.Error())
		// based on the error, we can return a specific error code
		// for example, if the error is a NotFound error, we can return a NotFound error code
		if errors.Is(err, apis.ErrNotFound) {
			return nil, errors.Status(codes.NotFound, err.Error(), &errdetails.ErrorInfo{
				Reason: err.Error(),
			})
		}
		return nil, errors.Status(codes.Internal, err.Error(), &errdetails.ErrorInfo{
			Reason: err.Error(),
		})
	}
	return &pb.ListModelsResponse{
		Models: intoPbModels(models),
	}, nil
}

func intoPbModels(models []apis.Model) []*pb.Model {
	var pbModels []*pb.Model
	for _, m := range models {
		pbModels = append(pbModels, &pb.Model{
			Id:   m.ID,
			Name: m.Name,
		})
	}
	return pbModels
}

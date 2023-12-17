package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/toochow-organization/bego/base/log"
	"github.com/toochow-organization/bego/external/apis"
	mock "github.com/toochow-organization/bego/external/apis/mock"
	pb "github.com/toochow-organization/bego/protocol/external/apis/v1"
)

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	ctx     context.Context
	server  Server
	service apis.ModelService
	storage *mock.ModelStorage
}

// SetupSuite initializes the basic context for testing.
func (s *ServerTestSuite) SetupSuite() {
	s.ctx = context.Background()

}

// SetupTest initializes the testing structure.
func (s *ServerTestSuite) SetupTest() {
	s.storage = mock.NewModelStorage(s.T())
	s.service = apis.NewModelService(s.storage)
	s.server = newHandler(log.L(), s.service)
}

func (s *ServerTestSuite) AfterTest() {
	// TODO: Add any cleanup logic here.
	//
	// usually this is where you would call s.storage.AssertExpectations(s.T())
	// where s.storage is an instance of a mock storage interface (using mockery).
	s.storage.AssertExpectations(s.T())
}

func (s *ServerTestSuite) TestGetVersion() {
	tests := []struct {
		name  string
		req   *pb.ListModelsRequest
		setup func()
		res   *pb.ListModelsResponse
		err   error
	}{
		{
			name: "success",
			req:  &pb.ListModelsRequest{},
			setup: func() {
				s.storage.On("ListModels", s.ctx).Return([]apis.Model{
					{Name: "version", ID: "123"},
				}, nil)
			},
			res: &pb.ListModelsResponse{Models: []*pb.Model{
				{Name: "version", Id: "123"},
			}},
			err: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setup()
			res, err := s.server.ListModels(s.ctx, tt.req)
			s.Equal(tt.err, err)
			s.Equal(tt.res, res)
		})
	}
}

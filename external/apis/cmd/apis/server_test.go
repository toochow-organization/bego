package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
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
	service apis.VersionService
	storage *mock.VersionStorage
}

// SetupSuite initializes the basic context for testing.
func (s *ServerTestSuite) SetupSuite() {
	s.ctx = context.Background()

}

// SetupTest initializes the testing structure.
func (s *ServerTestSuite) SetupTest() {
	s.storage = mock.NewVersionStorage(s.T())
	s.service = apis.NewVersionService(s.storage)
	s.server = newHandler(s.service)
}

func (s *ServerTestSuite) AfterTest() {
	// TODO: Add any cleanup logic here.
	//
	// usually this is where you would call s.storage.AssertExpectations(s.T())
	// where s.storage is an instance of a mock storage interface (using mockery).
}

func (s *ServerTestSuite) TestGetVersion() {
	tests := []struct {
		name  string
		req   *pb.GetVersionRequest
		setup func()
		res   *pb.GetVersionResponse
		err   error
	}{
		{
			name: "success",
			req:  &pb.GetVersionRequest{},
			setup: func() {
				s.storage.On("GetVersion", s.ctx).Return("1.0.1")
			},
			res: &pb.GetVersionResponse{Version: "1.0.1"},
			err: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setup()
			res, err := s.server.GetVersion(s.ctx, tt.req)
			s.Equal(tt.err, err)
			s.Equal(tt.res, res)
		})
	}
}

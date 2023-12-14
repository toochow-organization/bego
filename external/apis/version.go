package apis

import "context"

type VersionService struct {
	storage VersionStorage
}

// VersionStorage
//
// VersionStorage is the interface for the storage layer
// this is typically implemented by the dependency layer (which can be inmem, mysql, postgres, etc)
//
//go:generate mockery --name VersionStorage --output mock --outpkg mock --with-expecter
type VersionStorage interface {
	GetVersion(ctx context.Context) string
}

func NewVersionService(storage VersionStorage) VersionService {
	return VersionService{storage: storage}
}

// GetVersion
//
// Business logic for GetVersion
func (s *VersionService) GetVersion(ctx context.Context) string {
	return s.storage.GetVersion(ctx)
}

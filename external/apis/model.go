package apis

import "context"

type ModelService struct {
	storage ModelStorage
}

// ModelStorage
//
// ModelStorage is the interface for the storage layer
// this is typically implemented by the dependency layer (which can be inmem, mysql, postgres, etc)
//
//go:generate mockery --name ModelStorage --output mock --outpkg mock --with-expecter
type ModelStorage interface {
	ListModels(ctx context.Context) ([]Model, error)
}

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewModelService(storage ModelStorage) ModelService {
	return ModelService{storage: storage}
}

// ListModels
//
// Business logic for ListModels
func (s *ModelService) ListModels(ctx context.Context) ([]Model, error) {
	return s.storage.ListModels(ctx)
}

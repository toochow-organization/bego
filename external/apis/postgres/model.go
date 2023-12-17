package postgres

import (
	"context"

	log "github.com/toochow-organization/bego/base/log"
	"github.com/toochow-organization/bego/external/apis"
	"gorm.io/gorm"
)

var _ apis.ModelStorage = &ModelStorage{}

// struct of the model table
type model struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type ModelStorage struct {
	l *log.Logger
	// Gorm DB
	db *gorm.DB
	// Migrate the model table
	migrate bool
}

func ShouldMigrateModelStorage(doMigrate bool) func(*ModelStorage) {
	return func(s *ModelStorage) {
		s.migrate = doMigrate
	}
}

func NewModelStorage(l *log.Logger, db *gorm.DB, options ...func(*ModelStorage)) *ModelStorage {
	storage := &ModelStorage{
		l:  l,
		db: db,
	}
	for _, o := range options {
		o(storage)
	}
	if storage.migrate {
		storage.db.AutoMigrate(&model{})
	}
	return storage
}

// ListModels
//
// Returns a list of models.
func (s *ModelStorage) ListModels(ctx context.Context) ([]apis.Model, error) {
	var models []model
	if err := s.db.Find(&models).Error; err != nil {
		// TODO: Add more error handling.
		return nil, err
	}
	var result []apis.Model
	for _, m := range models {
		result = append(result, apis.Model{
			ID:   m.ID,
			Name: m.Name,
		})
	}
	return result, nil
}

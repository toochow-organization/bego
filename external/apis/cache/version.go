package inmem

import (
	"context"

	log "github.com/toochow-organization/bego/base/log"
	apis "github.com/toochow-organization/bego/external/apis"
)

var _ apis.VersionStorage = &VersionStorage{}

type VersionStorage struct {
	l *log.Logger
}

func NewVersionStorage(l *log.Logger) *VersionStorage {
	return &VersionStorage{
		l: l,
	}
}

// GetVersion
//
// GetVersion returns the version of the API
func (s *VersionStorage) GetVersion(ctx context.Context) string {
	s.l.Info(ctx, "GetVersion")
	return "1.0.0"
}

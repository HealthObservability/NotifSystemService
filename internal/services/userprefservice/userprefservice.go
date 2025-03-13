package userprefservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type UserPreferences interface {
	GetPreferences(ctx context.Context) (*UserPreferences, error)
	InsertPreference(ctx context.Context, preferences domains.UserPreferences) error
	DeletePreference(ctx context.Context, id int64) error
}

type service struct {
	adapters *adapters.Adapters
}

func (s service) GetPreferences(ctx context.Context) (*UserPreferences, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) InsertPreference(ctx context.Context, preferences domains.UserPreferences) error {
	//TODO implement me
	panic("implement me")
}

func (s service) DeletePreference(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func New(a *adapters.Adapters) UserPreferences {
	return service{
		adapters: a,
	}
}

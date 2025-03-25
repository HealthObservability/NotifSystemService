package userprefsvc

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
)

type UserPref interface {
	GetPreferences(ctx context.Context, filter domains.UserPrefs) ([]domains.UserPrefs, error)
	InsertPreference(ctx context.Context, preferences domains.UserPrefs) error
	DeletePreference(ctx context.Context, id int64) error
}

type service struct {
	Db ports.Database
}

func (s service) GetPreferences(ctx context.Context, filter domains.UserPrefs) ([]domains.UserPrefs, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) InsertPreference(ctx context.Context, preferences domains.UserPrefs) error {
	//TODO implement me
	panic("implement me")
}

func (s service) DeletePreference(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func New(db ports.Database) UserPref {
	return service{
		Db: db,
	}
}

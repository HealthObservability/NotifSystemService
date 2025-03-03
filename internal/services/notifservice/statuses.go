package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (s *Service) ChangeStatesStatus(ctx context.Context, n []domains.NotifState, newStatus string) error {
	length := len(n)
	ids := make([]uint64, length)

	for i := 0; i < length; i++ {
		ids = append(ids, n[i].ID)
	}

	if err := s.storage.SetNewStatuses(ctx, newStatus, ids); err != nil {
		return err
	}

	return nil
}

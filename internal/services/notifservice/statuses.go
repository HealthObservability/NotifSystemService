package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (s *service) UpdateSendStatus(ctx context.Context, n []domains.NotifState, status string) error {
	length := len(n)
	ids := make([]int64, length)

	for i := 0; i < length; i++ {
		ids[i] = n[i].ID
	}

	if err := s.Db.UpdateStatuses(ctx, status, ids); err != nil {
		return err
	}

	return nil
}

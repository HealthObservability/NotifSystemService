package notifservice

import "context"

func (s *Service) DeleteNotif(ctx context.Context, notifID int64) error {
	return s.storage.DeleteNotif(ctx, notifID)
}

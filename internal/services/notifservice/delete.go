package notifservice

import "context"

func (s *service) DeleteNotif(ctx context.Context, notifID int64) error {
	return s.Db.DeleteNotif(ctx, notifID)
}

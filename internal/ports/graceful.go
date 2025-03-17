package ports

import "context"

type GracefulShutdown interface {
	Shutdown(ctx context.Context) error
}

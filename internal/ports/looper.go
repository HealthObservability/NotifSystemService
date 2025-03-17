package ports

import "context"

type Looper interface {
	Start(ctx context.Context, handler func(ctx context.Context))
}

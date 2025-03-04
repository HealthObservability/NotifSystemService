package httpserver

import "context"

func (a *HTTPServAdapter) Shutdown(ctx context.Context) error {
	return a.mainServer.Shutdown(ctx)
}

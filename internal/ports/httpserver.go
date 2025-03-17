package ports

import "github.com/gin-gonic/gin"

type HTTPServer interface {
	GracefulShutdown
	SetHandlers(handler HTTPHandler)
	MustRun()
}

type HTTPHandler interface {
	AddNotification(c *gin.Context)
	DeleteNotification(c *gin.Context)
	GetNotifs(c *gin.Context)
}

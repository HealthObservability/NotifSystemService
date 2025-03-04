package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const httpTimeout = 10 * time.Second

type HTTPServAdapter struct {
	origins    []string
	mainServer *http.Server
	port       string
}

type HTTPHandler interface {
	AddNotification(c *gin.Context)
}

func (a *HTTPServAdapter) SetHandlers(handler HTTPHandler) {
	gin.SetMode(gin.ReleaseMode)
	mainRouter := gin.New()
	mainRouter.POST("/notifications", handler.AddNotification)
	//mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if len(a.origins) > 0 {
		mainRouter.Use(cors.New(cors.Config{
			AllowOrigins: a.origins,
		}))
	}

	a.mainServer = &http.Server{
		Addr:         ":" + a.port,
		Handler:      mainRouter,
		WriteTimeout: httpTimeout,
		ReadTimeout:  httpTimeout,
	}
}

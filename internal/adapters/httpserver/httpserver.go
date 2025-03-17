package httpserver

import (
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
	"net/http"
	"time"

	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const httpTimeout = 10 * time.Second

type HTTPServAdapter struct {
	origins    []string
	mainServer *http.Server
	port       string
}

func New(cfg config.HTTPServer) *HTTPServAdapter {
	return &HTTPServAdapter{
		origins: cfg.AllowOrigins,
		port:    cfg.Port,
	}
}

func (a *HTTPServAdapter) SetHandlers(handler ports.HTTPHandler) {
	gin.SetMode(gin.ReleaseMode)
	mainRouter := gin.New()
	mainRouter.POST("/notifications", handler.AddNotification)
	mainRouter.DELETE("/notifications/:id", handler.DeleteNotification)
	mainRouter.GET("/notifications/:id", handler.GetNotifs)
	mainRouter.GET("/notifications", handler.GetNotifs)
	// mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

func (a *HTTPServAdapter) MustRun() {
	log := logger.GetLogger().WithField("op", "HTTPServer.MustRun")
	log.WithField("port", a.port).Info("notif system service http server started")

	err := a.mainServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Fatal("metrics http server error")
	} else if err != nil {
		log.Info("metrics http server closed")
	}
}

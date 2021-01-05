package server

import (
	"context"
	"net/http"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"go.uber.org/fx"

	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/configurer"
	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/logger"
)

type server struct {
	cfg    *configurer.ServerCfg
	logger logger.Logger
	server *http.Server
}

func Run(lc fx.Lifecycle, cfg *configurer.ServerCfg, logger logger.Logger) {
	srv := server{
		cfg: cfg,
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			ReadTimeout:  cfg.Timeout.Duration,
			WriteTimeout: cfg.Timeout.Duration,
		},
		logger: logger.Named("gin_server"),
	}
	srv.server.Handler = srv.newRouter(cfg)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			srv.logger.Info("start listening at " + srv.cfg.Port)
			go func() {
				if err := srv.server.ListenAndServe(); err != nil {
					srv.logger.Fatal(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.logger.Info("wait for shutdown...")
			srv.logger.Sync()
			time.Sleep(2 * time.Second)
			return srv.server.Shutdown(ctx)
		},
	})
}

func (s *server) newRouter(cfg *configurer.ServerCfg) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(timeoutMiddleware(cfg.Timeout.Duration))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	ginpprof.Wrap(r)

	v1 := r.Group("/api/v1")
	for _, h := range v1handlers {
		v1.Handle(h.Method, h.Path, func(c *gin.Context) {
			h.HandlerFunc(c, s.logger)
		})
	}
	return r
}

func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

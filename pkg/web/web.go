package web

import (
	"app_blade/pkg/app"
	"app_blade/pkg/config"
	"app_blade/pkg/logging"
	"app_blade/pkg/profile"
	"app_blade/pkg/web/middleware"
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ app.AppRunner = (*WebServer)(nil)

type WebServer struct {
	*Config
	conf             config.Provider
	logger           logging.Provider
	httpServer       *http.Server
	initializeRouter InitializeRouter
}

// OnShutdown implements app.AppRunner.
func (ws *WebServer) OnShutdown(ctx context.Context) error {
	ws.logger.Infof("shutting down http server")
	if err := ws.httpServer.Shutdown(ctx); err != nil {
		ws.logger.Errorf("http shutdown error: %v", err)
	}
	return nil
}

// OnStart implements app.AppRunner.
func (ws *WebServer) OnStart(ctx context.Context) error {
	ws.logger.Infof("http start with addr: %s", ws.Addr())
	go func() {
		if err := ws.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			ws.logger.Errorf("http listen and serve error: %v", err)
		}
	}()
	return nil
}

func New(options ...Option) *WebServer {
	serve := &WebServer{
		Config:           DefaultConfig(),
		initializeRouter: func(e *gin.Engine) {},
	}

	for _, apply := range options {
		apply(serve)
	}

	if profile.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	route := gin.New()

	route.Use(cors.New(cors.Config{
		AllowMethods:     serve.Cors.AllowMethods,
		AllowOrigins:     serve.Cors.AllowOrigins,
		AllowHeaders:     serve.Cors.AllowHeaders,
		AllowCredentials: serve.Cors.AllowCredentials,
		MaxAge:           serve.Cors.MaxAge,
		AllowAllOrigins:  true,
	}))

	route.Use(middleware.NewRequestId())
	route.Use(middleware.NewLoggingRecovery())
	route.Use(middleware.NewLoggingRequest(serve.IgnorePatterns...))

	if serve.EnableDump {
		pprof.Register(route)
	}

	route.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// 初始化路由
	serve.initializeRouter(route)

	serve.httpServer = &http.Server{
		Addr:           serve.Addr(),
		ReadTimeout:    serve.ReadTimeout,
		WriteTimeout:   serve.WriteTimeout,
		MaxHeaderBytes: serve.MaxHeaderBytes,
		Handler:        route,
	}

	return serve
}

var ProviderSet = wire.NewSet(NewOptions, New)

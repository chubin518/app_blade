package web

import (
	"app_blade/pkg/config"
	"app_blade/pkg/logging"
	"time"
)

type Option func(*WebServer)

func WithHost(host string) Option {
	return func(s *WebServer) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *WebServer) {
		s.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *WebServer) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *WebServer) {
		s.WriteTimeout = timeout
	}
}

func WithMaxHeaderBytes(max int) Option {
	return func(s *WebServer) {
		s.MaxHeaderBytes = max
	}
}

func WithCors(cfg *CorsConfig) Option {
	return func(s *WebServer) {
		s.Cors = cfg
	}
}

func WithEnableDump(enable bool) Option {
	return func(s *WebServer) {
		s.EnableDump = enable
	}
}

func WithConfig(conf config.Provider) Option {
	return func(hs *WebServer) {
		hs.conf = conf
	}
}

func WithLogger(logger logging.Provider) Option {
	return func(hs *WebServer) {
		hs.logger = logger
	}
}

func WithInitializeRouter(initializeRouter InitializeRouter) Option {
	return func(s *WebServer) {
		s.initializeRouter = initializeRouter
	}
}

func NewOptions(conf config.Provider, logger logging.Provider, initializeRouter InitializeRouter) ([]Option, error) {
	cfg := DefaultConfig()
	if err := conf.UnmarshalKey("http", cfg); err != nil {
		return nil, err
	}
	return []Option{
		WithConfig(conf),
		WithLogger(logger),
		WithHost(cfg.Host),
		WithPort(cfg.Port),
		WithReadTimeout(cfg.ReadTimeout),
		WithWriteTimeout(cfg.WriteTimeout),
		WithMaxHeaderBytes(cfg.MaxHeaderBytes),
		WithCors(cfg.Cors),
		WithInitializeRouter(initializeRouter),
	}, nil
}

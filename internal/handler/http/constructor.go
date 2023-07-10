package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/http2"

	"github.com/kingmidas74/gonesis-engine/internal/env"
	"github.com/kingmidas74/gonesis-engine/internal/handler/http/middleware/no_cache"
)

type Params struct {
	fx.In

	Env *env.Env

	NoCacheMiddleware no_cache.Middleware
}

type Server struct {
	*http.Server
}

func NewServer(params Params) *Server {
	cwd, err := os.Getwd()
	if err != nil {
		zap.L().Error("Failed to get current working directory", zap.Error(err))
		return nil
	}

	assetsDir := filepath.Join(cwd, params.Env.Host.StaticFolder)

	r := chi.NewRouter()

	r.Use(params.NoCacheMiddleware.NoCache)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/", http.FileServer(http.Dir(assetsDir))).ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:         strings.Join([]string{":", params.Env.Host.Port}, ""),
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := http2.ConfigureServer(srv, &http2.Server{}); err != nil {
		zap.L().Error("Failed to configure HTTP/2 server", zap.Error(err))
		return nil
	}

	return &Server{srv}
}

func (srv *Server) Start() {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Could not listen on", zap.String("addr", srv.Addr), zap.Error(err))
		}
	}()

	zap.L().Info("Server is ready to handle requests", zap.String("addr", srv.Addr))
}

func (srv *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Could not gracefully shutdown the server", zap.Error(err))
	}

	zap.L().Info("Server stopped")
}

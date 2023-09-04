package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Amore14rn/888Starz/internal/config"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type App struct {
	cfg *config.Config

	router     *gin.Engine
	httpServer *http.Server
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})
	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"IP":   a.cfg.HTTP.IP,
		"Port": a.cfg.HTTP.Port,
	})
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	handler := a.router

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}
	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warning("server shutdown")
		default:
			logger.Fatal(err)
		}
	}

	httpErrChan := make(chan error, 1)
	httpShutdownChan := make(chan struct{})

	graceful.PerformGracefulShutdown(a.httpServer, httpErrChan, httpShutdownChan)

	return err
}

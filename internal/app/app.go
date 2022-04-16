package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eternalsad/user-auth-service/config"
	"github.com/eternalsad/user-auth-service/internal/handler"
	"github.com/eternalsad/user-auth-service/internal/repository"
	"github.com/eternalsad/user-auth-service/internal/service"
	"go.uber.org/zap"
)

func Run() error {
	cfg, err := config.NewConfiguration()
	if err != nil {
		return fmt.Errorf("error during parsing config file: %w", err)
	}
	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		return err
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer func() {
		err = logger.Sync()
	}()
	sugar := logger.Sugar()

	repo := repository.NewAuthRepo(db, cfg, sugar)
	service := service.NewService(repo, cfg, sugar)
	handler := handler.NewHandler(service, cfg, sugar)

	srv := &http.Server{
		Handler: handler.InitRoutes(),
	}
	errChan := make(chan error, 1)
	go func() {
		sugar.Infof("starting server on: %s", cfg.App.AppPort)
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-errChan:
	case <-quit:
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.App.AppShutDownTime)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Errorf("server was forced to shutdown: %s", err)
		return err
	}
	return nil
}

package main

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	_ "github.com/MaksimTP/CurrencyService/docs"

	"github.com/MaksimTP/CurrencyService/config"
	"github.com/MaksimTP/CurrencyService/internal/app"
	"github.com/sirupsen/logrus"
)

// @title Тестовое задание
// @version 0.0.1
// @description

// @host localhost:8080
// @BasePath /

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/

func main() {
	err := func(ctx context.Context) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		cfg, err := config.New()
		if err != nil {
			logrus.Fatalf("error while loading config: %s", err)
		}

		logrus.Info("loaded config..")

		go func() {
			defer cancel()
			app.Run(ctx, cfg)
		}()

		notifyCtx, notify := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
		defer notify()

		go func() {
			defer cancel()

			<-notifyCtx.Done()

			closer := make(chan struct{})

			go func() {
				closer <- struct{}{}
			}()

			shutdownCtx, shutdown := context.WithTimeout(context.Background(), time.Duration(cfg.HTTP.StartTimeout)*time.Second)
			defer shutdown()

			runtime.Gosched()

			select {
			case <-closer:
				logrus.Info("shutting down gracefully..")
			case <-shutdownCtx.Done():
				logrus.Error("shutting down forcefully..")
			}
		}()

		<-ctx.Done()

		cancel()

		return nil
	}(context.Background())
	if err != nil {
		logrus.Errorf("error running a server: %s", err)
		return
	}
}

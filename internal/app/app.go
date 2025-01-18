package app

import (
	"context"

	"github.com/MaksimTP/CurrencyService/config"
	v1 "github.com/MaksimTP/CurrencyService/internal/controller/http/api/v1"
	"github.com/MaksimTP/CurrencyService/internal/usecase"
	"github.com/MaksimTP/CurrencyService/internal/usecase/repo"
	"github.com/MaksimTP/CurrencyService/pkg/coinapi"
	"github.com/MaksimTP/CurrencyService/pkg/httpserver"
	"github.com/MaksimTP/CurrencyService/pkg/observe"
	"github.com/MaksimTP/CurrencyService/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context, cfg *config.Config) {
	if cfg == nil {
		logrus.Fatalf("config is nil, shutting down...")
	}

	pg, err := postgres.New(ctx, cfg.DBInfo())
	if err != nil {
		logrus.Fatalf("error while connecting to db: %s, params:%s", err.Error(), cfg.DBInfo())
	}
	defer pg.Close()

	logrus.Info("connected to db..")

	handler := gin.Default()

	r, err := repo.NewRepo(pg)
	if err != nil {
		logrus.Fatalf("error while creating repo %s", err)
	}

	cryptoSymbols, err := observe.GetCryptoSymbols("cryptosymbols.json")
	if err != nil {
		logrus.Fatalf("failed to fetch cryptosymbols %s", err)
	}

	coinAPI := coinapi.New(cfg.CoinAPI.Token)

	observer := observe.New(coinAPI, cfg.Observer.ObserveTime, cryptoSymbols)
	currencyUseCase, err := usecase.NewCurrencyUseCase(r, observer)
	if err != nil {
		logrus.Fatalf("error while setting up currency usecase %s", err)
	}

	err = v1.NewRouterCurrency(handler, currencyUseCase)
	if err != nil {
		logrus.Fatalf("error while setting up currency router %s", err)
	}

	err = v1.NewSwagger(handler)
	if err != nil {
		logrus.Fatalf("error while setting up swagger router %s", err)
	}

	server := httpserver.New()

	observer.ObserveCurrency("BTC")
	go currencyUseCase.Observe()

	logrus.Infof("starting server on port %s", cfg.HTTP.Port)
	if err := server.Run(
		cfg.HTTP.Port,
		cfg.HTTP.MaxHeaderBytes,
		cfg.HTTP.ReadTimeout,
		cfg.HTTP.WriteTimeout,
		handler,
	); err != nil {
		logrus.Fatalf("error while trying to serve on addr [%s], error:[%s]", cfg.HTTP.Port, err)
	}
}

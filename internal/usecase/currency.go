package usecase

import (
	"time"

	"github.com/MaksimTP/CurrencyService/internal/entity/request"
	"github.com/MaksimTP/CurrencyService/internal/entity/response"
	"github.com/MaksimTP/CurrencyService/pkg/err"
	"github.com/MaksimTP/CurrencyService/pkg/observe"
	"github.com/sirupsen/logrus"
)

type CurrencyRepository interface {
	GetCurrencyPrice(req request.GetCurrencyPrice) response.GetCurrencyPrice
	SaveCurrency(req request.SaveCurrency) response.SaveCurrency
}

type CurrencyUC struct {
	repo     CurrencyRepository
	observer *observe.Observer
}

func NewCurrencyUseCase(r CurrencyRepository, observer *observe.Observer) (*CurrencyUC, error) {
	if r == nil {
		return nil, err.ErrPostgresNil
	}

	if observer == nil {
		return nil, err.ErrNilParam
	}

	return &CurrencyUC{
		repo:     r,
		observer: observer,
	}, nil
}

func (uc *CurrencyUC) AddCurrency(req request.AddCurrency) response.AddCurrency {
	err := uc.observer.ObserveCurrency(req.Coin)

	resp := response.AddCurrency{
		Status: map[string]interface{}{
			"error": err,
		},
	}
	return resp
}
func (uc *CurrencyUC) DeleteCurrency(req request.DeleteCurrency) response.DeleteCurrency {
	err := uc.observer.StopObserveCurrency(req.Coin)

	resp := response.DeleteCurrency{
		Status: map[string]interface{}{
			"error": err,
		},
	}

	return resp
}

func (uc *CurrencyUC) Observe() {
	for {
		time.Sleep(time.Duration(uc.observer.CollectTime) * time.Second)

		if len(uc.observer.CurrenciesToObserve) == 0 {
			logrus.Warn("No currencies to observe")
		}

		for currency := range uc.observer.CurrenciesToObserve {

			go func() {
				price, err := uc.observer.CoinAPI.GetPrice(currency)
				if err != nil {
					logrus.Warnf("error while fetching price %s", err)
				}

				req := request.SaveCurrency{
					Coin:      price.Symbol,
					Price:     price.Price,
					Timestamp: price.Timestamp,
				}

				resp := uc.repo.SaveCurrency(req)

				if resp.Status["error"] != nil {
					logrus.Warnf("error while saving currency price to DB %s", err)
				}

				logrus.Infof("%s %f %d", price.Symbol, price.Price, price.Timestamp)
			}()
		}
	}
}

func (uc *CurrencyUC) GetCurrencyPrice(req request.GetCurrencyPrice) response.GetCurrencyPrice {
	req.Coin = req.Coin + "USD" // rework later
	return uc.repo.GetCurrencyPrice(req)
}

// package repo is a layer that represents interaction between usecase and database.
package repo

import (
	"github.com/MaksimTP/CurrencyService/internal/entity/request"
	"github.com/MaksimTP/CurrencyService/internal/entity/response"
	"github.com/MaksimTP/CurrencyService/pkg/err"
	"github.com/MaksimTP/CurrencyService/pkg/postgres"
)

type Repo struct {
	repo *postgres.Postgres
}

func NewRepo(repo *postgres.Postgres) (*Repo, error) {
	if repo == nil {
		return nil, err.ErrPostgresNil
	}
	return &Repo{
		repo: repo,
	}, nil
}

func (r *Repo) GetCurrencyPrice(req request.GetCurrencyPrice) response.GetCurrencyPrice {
	return r.repo.GetCurrencyPrice(req)
}

func (r *Repo) SaveCurrency(req request.SaveCurrency) response.SaveCurrency {
	return r.repo.SaveCurrency(req)
}

package postgres

import (
	"context"

	"github.com/MaksimTP/CurrencyService/internal/entity/request"
	"github.com/MaksimTP/CurrencyService/internal/entity/response"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func New(ctx context.Context, url string) (*Postgres, error) {
	conn, err := sqlx.ConnectContext(ctx, "postgres", url)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	// Disable "missing destination name" error
	conn = conn.Unsafe()

	return &Postgres{
		db: conn,
	}, nil
}

func (p *Postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *Postgres) GetCurrencyPrice(req request.GetCurrencyPrice) response.GetCurrencyPrice {
	resp := response.GetCurrencyPrice{
		Status: map[string]interface{}{},
	}

	err := p.db.Get(&resp, GetCurrencyPriceQuery, req.Coin, req.Timestamp)
	resp.Status["error"] = err
	return resp
}

func (p *Postgres) SaveCurrency(req request.SaveCurrency) response.SaveCurrency {
	_, err := p.db.NamedExec(SaveCurrencyQuery, req)
	return response.SaveCurrency{
		Status: map[string]interface{}{
			"error": err,
		},
	}
}

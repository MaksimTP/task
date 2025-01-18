package observe

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/MaksimTP/CurrencyService/pkg/coinapi"
	"github.com/MaksimTP/CurrencyService/pkg/err"
)

type Observer struct {
	CurrenciesToObserve map[string]struct{}
	CollectTime         int
	CoinAPI             *coinapi.CoinAPI

	cryptoSymbols map[string]struct{}

	mu sync.Mutex
}

func GetCryptoSymbols(filename string) (map[string]struct{}, error) {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	symbols := []string{}

	err = json.Unmarshal(fd, &symbols)
	if err != nil {
		return nil, err
	}

	cryptoSymbols := make(map[string]struct{})

	for _, symbol := range symbols {
		cryptoSymbols[symbol] = struct{}{}
	}

	return cryptoSymbols, nil
}

func New(coinAPI *coinapi.CoinAPI, collectTime int, cryptosymbols map[string]struct{}) *Observer {

	return &Observer{
		CurrenciesToObserve: map[string]struct{}{},
		CoinAPI:             coinAPI,
		CollectTime:         collectTime,
		cryptoSymbols:       cryptosymbols,
	}
}

func (o *Observer) ObserveCurrency(currency string) error {
	currency = currency + "USD" // rework later
	o.mu.Lock()
	defer o.mu.Unlock()

	if _, exists := o.cryptoSymbols[currency]; !exists {
		return err.ErrCurrencyNotFound
	}
	o.CurrenciesToObserve[currency] = struct{}{}
	return nil
}

func (o *Observer) StopObserveCurrency(currency string) error {
	currency = currency + "USD" // rework later
	o.mu.Lock()
	defer o.mu.Unlock()

	if _, exists := o.cryptoSymbols[currency]; !exists {
		return err.ErrCurrencyNotFound
	}

	delete(o.CurrenciesToObserve, currency)
	return nil
}

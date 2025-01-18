package coinapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoinAPI struct {
	token string
}

type Price struct {
	Symbol    string
	Price     float64 `json:",string"`
	Timestamp int64
}

//	https://api.api-ninjas.com/v1/cryptoprice?symbol=LTCBTC

func New(token string) *CoinAPI {
	return &CoinAPI{
		token: token,
	}
}

func (c *CoinAPI) GetPrice(currency string) (Price, error) {

	url := fmt.Sprintf("https://api.api-ninjas.com/v1/cryptoprice?symbol=%s", currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Price{}, err
	}

	req.Header.Set("X-Api-Key", c.token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return Price{}, err
	}
	defer resp.Body.Close()

	var price Price

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Price{}, err
	}
	err = json.Unmarshal(data, &price)
	return price, err
}

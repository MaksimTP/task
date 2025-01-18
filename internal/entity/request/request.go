package request

type AddCurrency struct {
	Coin string `json:"coin" db:"coin"`
}

type DeleteCurrency struct {
	Coin string `json:"coin" db:"coin"`
}

type GetCurrencyPrice struct {
	Coin      string `json:"coin" db:"coin"`
	Timestamp int64  `json:"timestamp" db:"timestamp"`
}

type SaveCurrency struct {
	Coin      string  `json:"coin" db:"coin"`
	Price     float64 `json:"price" db:"price"`
	Timestamp int64   `json:"timestamp" db:"timestamp"`
}

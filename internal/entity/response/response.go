package response

type AddCurrency struct {
	Status map[string]interface{} `json:"status" db:"status"`
}

type DeleteCurrency struct {
	Status map[string]interface{} `json:"status" db:"status"`
}

type GetCurrencyPrice struct {
	Coin      string                 `json:"coin" db:"coin"`
	Price     float64                `json:"price" db:"price"`
	Timestamp int64                  `json:"timestamp" db:"timestamp"`
	Status    map[string]interface{} `json:"status" db:"status"`
}

type SaveCurrency struct {
	Status map[string]interface{} `json:"status" db:"status"`
}

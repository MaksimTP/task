package postgres

const (
	SaveCurrencyQuery     = "INSERT INTO price (coin, price, timestamp) VALUES (:coin, :price, :timestamp)"
	GetCurrencyPriceQuery = "SELECT * FROM price WHERE coin=$1 ORDER BY ABS(timestamp - $2) LIMIT 1"
)

package err

import "errors"

var ErrNilParam = errors.New("one or more parameters in function are nil")
var ErrPostgresNil = errors.New("postgres instance is nil")
var ErrCurrencyNotFound = errors.New("currency not found")

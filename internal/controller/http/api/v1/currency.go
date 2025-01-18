package v1

import (
	"net/http"
	"strconv"

	"github.com/MaksimTP/CurrencyService/internal/entity/request"
	"github.com/MaksimTP/CurrencyService/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Currency struct {
	t *usecase.CurrencyUC
}

func newCurrency(h *gin.RouterGroup, t *usecase.CurrencyUC) {
	r := &Currency{t}
	h.GET("/add/:coin", r.AddCurrency)
	h.GET("/remove/:coin", r.DeleteCurrency)
	h.GET("/price/:coin/:timestamp", r.GetCurrencyPrice)
}

// ShowAccount godoc
// @Summary adds currency for observing
// @Description adds currency for observing
// @Tags currency
// @Accept json
// @Produce json
// @Router /api/v1/currency/add/{coin} [get]
// @Param coin path string true "Coin Name"
// @Success 200 {object} response.AddCurrency
func (r *Currency) AddCurrency(c *gin.Context) {
	req := request.AddCurrency{
		Coin: c.Param("coin"),
	}

	resp := r.t.AddCurrency(req)

	c.JSON(http.StatusOK, resp)
}

// ShowAccount godoc
// @Summary delete currency from observing
// @Description get recommendations by user id
// @Tags currency
// @Accept json
// @Produce json
// @Router /api/v1/currency/remove/{coin} [get]
// @Param coin path string true "Coin Name"
// @Success 200 {object} response.DeleteCurrency
func (r *Currency) DeleteCurrency(c *gin.Context) {
	req := request.DeleteCurrency{
		Coin: c.Param("coin"),
	}

	resp := r.t.DeleteCurrency(req)

	c.JSON(http.StatusOK, resp)
}

// ShowAccount godoc
// @Summary get currency price
// @Description get currency price by timestamp
// @Tags currency
// @Accept json
// @Produce json
// @Router /api/v1/currency/price/{coin}/{timestamp} [get]
// @Param coin path string true "Coin Name"
// @Param timestamp path int false "Timestamp"
// @Success 200 {object} response.GetCurrencyPrice
func (r *Currency) GetCurrencyPrice(c *gin.Context) {
	coin := c.Param("coin")
	timestamp, err := strconv.Atoi(c.Param("timestamp"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	req := request.GetCurrencyPrice{
		Coin:      coin,
		Timestamp: int64(timestamp),
	}

	resp := r.t.GetCurrencyPrice(req)

	c.JSON(http.StatusOK, resp)
}

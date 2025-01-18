package v1

import (
	"net/http"

	"github.com/MaksimTP/CurrencyService/internal/usecase"
	"github.com/MaksimTP/CurrencyService/pkg/err"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouterCurrency(handler *gin.Engine, t *usecase.CurrencyUC) error {
	if handler == nil || t == nil {
		return err.ErrNilParam
	}

	h := handler.Group("/api/v1/currency")
	{
		newCurrency(h, t)
	}

	return nil
}

func NewSwagger(handler *gin.Engine) error {
	if handler == nil {
		return err.ErrNilParam
	}
	handler.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/swagger/index.html")
	})
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return nil
}

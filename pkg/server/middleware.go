package server

import (
	"net/http"

	"github.com/ShinyShield/AbyShopServ/pkg/entities"
	"github.com/ShinyShield/AbyShopServ/pkg/presenter"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
	"github.com/labstack/echo"
)

// token
func memberMiddleware(hfc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var result presenter.Result

		//  Header  token
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}

		account := new(entities.Account)

		//  token
		query := store.DB.Model(entities.Account{}).
			Where("token = ?", token).
			First(account)
		if query.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if query.Error != nil {
			result.Code = presenter.StatusServerError
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		ctx.Set("account", account)

		return hfc(ctx)
	}
}

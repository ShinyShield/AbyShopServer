package presenter

import (
	"net/http"

	"github.com/ShinyShield/AbyShopServ/pkg/entities"
	"github.com/ShinyShield/AbyShopServ/pkg/usecases/logout"
	"github.com/labstack/echo"
)

func Logout(ctx echo.Context) error {

	var result Result

	account := ctx.Get("account").(*entities.Account)

	var input logout.Input
	input.AccountID = account.ID

	if err := logout.Exec(input); err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}

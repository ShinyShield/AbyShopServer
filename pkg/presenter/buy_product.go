package presenter

import (
	"net/http"

	"github.com/ShinyShield/AbyShopServer/pkg/entities"
	update "github.com/ShinyShield/AbyShopServer/pkg/usecases/buy_product"
	"github.com/labstack/echo"
)

type BuyProductParam struct {
	Product_ID int32
	Number     int32
}

func BuyProduct(ctx echo.Context) error {

	var param BuyProductParam
	var result Result

	if err := ctx.Bind(&param); err != nil {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	account := ctx.Get("account").(*entities.Account)

	var input update.Input
	input.CustomerID = account.ID
	input.Number = param.Number
	input.ProductID = param.Product_ID

	err := update.Exec(input)
	if err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}

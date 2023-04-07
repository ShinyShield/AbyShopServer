package presenter

import (
	"net/http"

	"github.com/ShinyShield/AbyShopServ/pkg/usecases/search"
	"github.com/labstack/echo"
)

type SearchParam struct {
	Query string `query:"q"`
}

func Search(ctx echo.Context) error {

	var result Result
	var param SearchParam

	if err := ctx.Bind(&param); err != nil {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input search.Input
	input.Query = param.Query

	output, err := search.Exec(input)
	if err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess
	result.Count = output.Count
	result.Data = output.Products

	return ctx.JSON(http.StatusOK, result)
}

package presenter

import (
	"net/http"
	"strconv"

	query "github.com/ShinyShield/AbyShopServ/pkg/usecases/query_product"
	"github.com/labstack/echo"
)

// 搜尋某產品關鍵字
func QueryProduct(ctx echo.Context) error {

	var result Result

	ProductID := ctx.Param("product_id")
	if ProductID == "" {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input query.Input

	int32ID, err := strconv.ParseInt(ProductID, 10, 32)
	if err != nil {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	input.ProductID = int32(int32ID)

	output, err := query.Exec(input)
	if err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess
	result.Data = output.Product

	return ctx.JSON(http.StatusOK, result)
}

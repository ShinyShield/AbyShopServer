package presenter

import (
	"net/http"

	main "github.com/ShinyShield/AbyShopServer/pkg/usecases/main_page"
	"github.com/labstack/echo"
)

func MainPage(ctx echo.Context) error {

	var result Result

	output, err := main.Exec()
	if err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess
	result.Data = output

	return ctx.JSON(http.StatusOK, result)
}

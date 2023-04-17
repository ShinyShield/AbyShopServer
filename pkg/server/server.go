package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {

	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		middleware.Logger(),
	)

	generalRoute(e)
	memberRoute(e, memberMiddleware)

	e.Logger.Fatal(e.Start(":1323"))
}

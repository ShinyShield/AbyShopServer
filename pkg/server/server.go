package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {

	e := echo.New()
	// 註冊中間件
	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		middleware.Logger(),
	)
	// 註冊各個 route
	generalRoute(e)
	memberRoute(e, memberMiddleware)

	e.Logger.Fatal(e.Start(":1323"))
}

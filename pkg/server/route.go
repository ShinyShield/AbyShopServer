package server

import (
	"github.com/ShinyShield/AbyShopServ/pkg/presenter"
	"github.com/labstack/echo"
)

// route
func generalRoute(e *echo.Echo) {
	e.GET("/main", presenter.MainPage)
	e.GET("/search", presenter.Search)
	e.GET("/query/product/:product_id", presenter.QueryProduct)

	e.POST("/login", presenter.Login)
}

// route
func memberRoute(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	routers := e.Group("/member", middlewares...)

	routers.POST("/buy", presenter.BuyProduct)
	routers.DELETE("/logout", presenter.Logout)
}

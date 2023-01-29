package core

import (
	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/wire"
)

func InitRoutes(e *echo.Echo) {
	{
		cart := e.Group("/cart")
		cartController := wire.CartController()
		cartController.Mount(cart)
	}
}

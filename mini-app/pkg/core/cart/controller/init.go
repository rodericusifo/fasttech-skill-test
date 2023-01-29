package controller_cart

import (
	"github.com/labstack/echo/v4"
	service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service"
)

type CartController struct {
	CartService service_cart.ICartService
}

func InitCartController(cartService service_cart.ICartService) *CartController {
	return &CartController{CartService: cartService}
}

func (cartController *CartController) Mount(group *echo.Group) {
	group.POST("/add-product", cartController.AddProductToCart)
	group.GET("/list", cartController.GetListCart)
	group.DELETE("/delete-product", cartController.DeleteProductFromCart)
}

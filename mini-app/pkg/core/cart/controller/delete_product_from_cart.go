package controller_cart

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/response"
	request_controller_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/controller/request"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/validator"
)

func (c *CartController) DeleteProductFromCart(ctx echo.Context) error {
	reqBody := new(request_controller_cart.DeleteProductFromCartRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.CartService.DeleteProductFromCart(&dto_service_cart.DeleteProductFromCartDTO{
		ProductCode: reqBody.ProductCode,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess[any]("success delete product from cart", nil))
}

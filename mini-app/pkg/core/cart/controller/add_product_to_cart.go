package controller_cart

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/response"
	request_controller_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/controller/request"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/validator"
)

func (c *CartController) AddProductToCart(ctx echo.Context) error {
	reqBody := new(request_controller_cart.AddProductToCartRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.CartService.AddProductToCart(&dto_service_cart.AddProductToCartDTO{
		ProductName: reqBody.ProductName,
		ProductCode: reqBody.ProductCode,
		Quantity:    reqBody.Quantity,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.ResponseSuccess[any]("success add product to cart", nil))
}

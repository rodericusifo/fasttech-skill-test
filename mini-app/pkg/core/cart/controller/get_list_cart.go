package controller_cart

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/response"
	request_controller_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/controller/request"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/validator"
)

func (c *CartController) GetListCart(ctx echo.Context) error {
	reqQuery := new(request_controller_cart.GetListCartRequestQuery)
	if err := validator.ValidateRequest(ctx, reqQuery); err != nil {
		return err
	}

	cartDTOGetListResponse, err := c.CartService.GetListCart(&dto_service_cart.GetListCartPayloadDTO{
		ProductName: reqQuery.ProductName,
		Quantity:    reqQuery.Quantity,
	})
	if err != nil {
		return err
	}

	res := make([]string, 0)
	for _, cartDTO := range cartDTOGetListResponse {
		res = append(res, fmt.Sprintf("[ %s ] - [ %s ] - [ %d ]", cartDTO.ProductCode, cartDTO.ProductName, cartDTO.Quantity))
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("success get list cart", res))
}

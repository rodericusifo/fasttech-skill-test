package service_cart

import (
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (s *CartService) GetListCart(payload *dto_service_cart.GetListCartPayloadDTO) ([]*dto_service_cart.GetListCartDTO, error) {
	cartModelGetListResult, err := s.CartResource.GetListCart(nil, &model.Cart{
		ProductName: payload.ProductName,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		return nil, err
	}

	cartDTOGetListResult := make([]*dto_service_cart.GetListCartDTO, 0)
	for _, cartModel := range cartModelGetListResult {
		cartDTOGetListResult = append(cartDTOGetListResult, &dto_service_cart.GetListCartDTO{
			ProductName: cartModel.ProductName,
			ProductCode: cartModel.ProductCode,
			Quantity:    cartModel.Quantity,
		})
	}

	return cartDTOGetListResult, nil
}

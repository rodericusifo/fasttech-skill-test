package service_cart

import (
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (s *CartService) DeleteProductFromCart(payload *dto_service_cart.DeleteProductFromCartDTO) error {
	cartModelGetResult, err := s.CartResource.GetCart(nil, &model.Cart{
		ProductCode: payload.ProductCode,
	})
	if err != nil {
		return err
	}

	err = s.CartResource.DeleteCart(cartModelGetResult)
	if err != nil {
		return err
	}

	return nil
}

package service_cart

import (
	resource_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/resource"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
)

type ICartService interface {
	DeleteProductFromCart(payload *dto_service_cart.DeleteProductFromCartDTO) error
	AddProductToCart(payload *dto_service_cart.AddProductToCartDTO) error
	GetListCart(payload *dto_service_cart.GetListCartPayloadDTO) ([]*dto_service_cart.GetListCartDTO, error)
}

type CartService struct {
	CartResource resource_cart.ICartResource
}

func InitCartService(cartResource resource_cart.ICartResource) ICartService {
	return &CartService{
		CartResource: cartResource,
	}
}

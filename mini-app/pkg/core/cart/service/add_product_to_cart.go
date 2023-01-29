package service_cart

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	"gorm.io/gorm"
)

type (
	AddProductToCartDTOToCart struct {
		ProductCode string `json:"product_code"`
		ProductName string `json:"product_name"`
		Quantity    int64  `json:"quantity"`
	}
)

func (s *CartService) AddProductToCart(payload *dto_service_cart.AddProductToCartDTO) error {
	cartModelGetResult, err := s.CartResource.GetCart(nil, &model.Cart{
		ProductCode: payload.ProductCode,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if cartModelGetResult != nil {
		cartModelGetResult.Quantity += payload.Quantity

		err = s.CartResource.UpdateCart(cartModelGetResult)

		if err != nil {
			return err
		}

		return nil
	}

	cartModel, err := util.ConvertToStruct[*model.Cart](
		&AddProductToCartDTOToCart{
			ProductCode: payload.ProductCode,
			ProductName: payload.ProductName,
			Quantity: payload.Quantity,
		},
	)
	if err != nil {
		return err
	}

	err = s.CartResource.CreateCart(cartModel)
	if err != nil {
		return err
	}

	return nil
}

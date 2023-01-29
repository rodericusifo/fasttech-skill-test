package resource_cart

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartResource) CreateCart(payload *model.Cart) error {
	return r.Postgres.CartRepository.CreateCart(payload)
}

package resource_cart

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartResource) DeleteCart(payload *model.Cart) error {
	return r.Postgres.CartRepository.DeleteCart(payload)
}

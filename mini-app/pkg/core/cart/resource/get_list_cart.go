package resource_cart

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartResource) GetListCart(query *types.Query, payload *model.Cart) ([]*model.Cart, error) {
	return r.Postgres.CartRepository.GetListCart(query, payload)
}

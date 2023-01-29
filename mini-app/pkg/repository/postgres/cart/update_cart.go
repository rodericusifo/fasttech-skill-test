package cart_postgres_repository

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartRepository) UpdateCart(payload *model.Cart) error {
	cart := new(model.Cart)

	sql := r.db.Model(cart)

	if payload != nil {
		cart = payload
	}

	if err := sql.Where("id = ?", payload.ID).Save(cart).Error; err != nil {
		return err
	}

	return nil
}

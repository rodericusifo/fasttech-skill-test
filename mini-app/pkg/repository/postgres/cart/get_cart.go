package cart_postgres_repository

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartRepository) GetCart(query *types.Query, payload *model.Cart) (*model.Cart, error) {
	cart := new(model.Cart)

	sql := r.db.Model(cart)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(util.MergeSlices(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.ProductCode != "" {
			sql = sql.Where("product_code = ?", payload.ProductCode)
		}
	}

	if err := sql.First(cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

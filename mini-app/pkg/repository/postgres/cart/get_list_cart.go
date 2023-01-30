package cart_postgres_repository

import (
	"fmt"

	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

func (r *CartRepository) GetListCart(query *types.Query, payload *model.Cart) ([]*model.Cart, error){
	cart := new(model.Cart)
	carts := make([]*model.Cart, 0)

	sql := r.db.Model(cart)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(util.MergeSlices(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.ProductName != "" {
			sql = sql.Where("product_name LIKE ?", fmt.Sprintf("%%%s%%", payload.ProductName))
		}
		if payload.Quantity != 0 {
			sql = sql.Where("quantity = ?", payload.Quantity)
		}
	}

	if err := sql.Find(&carts).Error; err != nil {
		return nil, err
	}

	return carts, nil
}
package cart_postgres_repository

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(payload *model.Cart) error
	UpdateCart(payload *model.Cart) error
	DeleteCart(payload *model.Cart) error
	GetCart(query *types.Query, payload *model.Cart) (*model.Cart, error)
	GetListCart(query *types.Query, payload *model.Cart) ([]*model.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func InitCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{
		db: db,
	}
}

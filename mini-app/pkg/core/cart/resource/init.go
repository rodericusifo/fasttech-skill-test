package resource_cart

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	postgres_repository "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/repository/postgres"
)

type ICartResource interface {
	DeleteCart(payload *model.Cart) error
	CreateCart(payload *model.Cart) error
	UpdateCart(payload *model.Cart) error
	GetCart(query *types.Query, payload *model.Cart) (*model.Cart, error)
	GetListCart(query *types.Query, payload *model.Cart) ([]*model.Cart, error)
}

type CartResource struct {
	Postgres *postgres_repository.PostgresRepository
}

func InitCartResource() ICartResource {
	var (
		postgres = postgres_repository.InitPostgresRepository()
	)

	return &CartResource{
		Postgres: postgres,
	}
}

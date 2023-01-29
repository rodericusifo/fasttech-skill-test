package postgres_repository

import (
	cart_postgres_repository "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/repository/postgres/cart"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/config"
)

type PostgresRepository struct {
	CartRepository cart_postgres_repository.ICartRepository
}

func InitPostgresRepository() *PostgresRepository {
	var (
		cartRepository = cart_postgres_repository.InitCartRepository(config.DB)
	)

	return &PostgresRepository{
		CartRepository: cartRepository,
	}
}

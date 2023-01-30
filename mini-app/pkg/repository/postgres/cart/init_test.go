package cart_postgres_repository

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	mockCartRepository ICartRepository
	mockQuery          sqlmock.Sqlmock
)

var (
	mockDate time.Time
	mockUUID string
)

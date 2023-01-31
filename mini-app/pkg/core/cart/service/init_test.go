package service_cart

import (
	"time"

	"github.com/rodericusifo/fasttech-skill-test/mini-app/mocks"
)

var (
	mockCartResource *mocks.ICartResource
	mockCartService ICartService
)

var (
	mockDate time.Time
	mockUUID string
)

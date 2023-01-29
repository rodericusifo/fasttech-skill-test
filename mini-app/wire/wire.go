//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	resource_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/resource"
	service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service"
	controller_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/controller"
)

func CartController() *controller_cart.CartController {
	wire.Build(
		controller_cart.InitCartController,
		service_cart.InitCartService,
		resource_cart.InitCartResource,
	)
	return &controller_cart.CartController {}
}
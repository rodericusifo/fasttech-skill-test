package migration

import (
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
)

var (
	AutoMigrateModelList = []interface{}{
		&model.Cart{},
	}
)
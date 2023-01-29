package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	libs_validator "github.com/rodericusifo/fasttech-skill-test/mini-app/libs/validator"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/config"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/custom"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/util"
)

func init() {
	config.ConfigApps()
	config.ConfigureDatabaseSQL(constant.POSTGRES)
}

func main() {
	e := echo.New()
	e.Validator = &libs_validator.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = custom.CustomHTTPErrorHandler

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	core.InitRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", util.GetPortApp())))
}

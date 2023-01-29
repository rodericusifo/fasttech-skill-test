package util

import (
	"strconv"

	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/config"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}

func GetEnvironmentType() constant.EnvironmentType {
	return config.Environment
}

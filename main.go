package main

import (
	"github.com/json-iterator/go/extra"
	"metis/config"
	"metis/router"
	"metis/util/logger"
	"strings"
	"time"
)

func lowerCamelCase(f string) string {
	return strings.ToLower(f[:1]) + f[1:]
}

func main() {
	useLogger := logger.UseLogger()
	tomlConfig := config.TomlConfig()
	baseRouter := router.BaseRouter()

	extra.RegisterFuzzyDecoders()
	extra.RegisterTimeAsInt64Codec(time.Millisecond)
	extra.SetNamingStrategy(lowerCamelCase)

	// docs.SwaggerInfo.BasePath = "/"
	// baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	useLogger.Sugar().Infof("config -> %v", tomlConfig)
	useLogger.Info("hello GTP, I'm born.")

	err := baseRouter.Run(":9999")
	if err != nil {
		return
	}
}

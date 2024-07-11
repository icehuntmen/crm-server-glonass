package main

import (
	"context"
	"crm-glonass/api"
	"crm-glonass/config"
	"crm-glonass/data/cache"
	mongox "crm-glonass/data/mongox"
	"crm-glonass/middlewares"
	"crm-glonass/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())
var version = middlewares.GetVersion()

func main() {

	ctx := context.TODO()

	logger.Info(logging.General, logging.StartUp, "Started server...", map[logging.ExtraKey]interface{}{"Version": version})

	conf := config.GetConfig()

	database := mongox.Connection(conf, ctx, logger)
	cache.InitRedis(conf, ctx)

	api.InitialServer(conf, database, logger)
}

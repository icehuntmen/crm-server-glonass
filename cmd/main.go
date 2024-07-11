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
var version, _ = middlewares.GetVersion()

func main() {

	ctx := context.TODO()

	logger.Info(logging.General, logging.StartUp, "Started server...", map[logging.ExtraKey]interface{}{"Version": version})

	conf := config.GetConfig()

	database, _ := mongox.Connection(conf, ctx, logger)
	cache.InitRedis(conf, ctx)
	logger.Infof("Listening on Swagger http://localhost:%d/swagger/index.html", conf.Server.IPort)
	api.InitialServer(conf, database, logger)

}

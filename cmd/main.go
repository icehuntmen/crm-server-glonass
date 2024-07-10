package main

import (
	"context"
	"crm-glonass/api"
	"crm-glonass/config"
	"crm-glonass/data/cache"
	mongox "crm-glonass/data/mongox"
	"crm-glonass/pkg/logging"
)

var logcod = logging.NewLogger(config.GetConfig())

func main() {

	ctx := context.TODO()

	logcod.Info(logging.General, logging.StartUp, "Started server...", nil)

	conf := config.GetConfig()

	database := mongox.Connection(conf, ctx, logcod)
	cache.InitRedis(conf, ctx)

	api.InitialServer(conf, database, logcod)
}

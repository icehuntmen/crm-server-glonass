package main

import (
	"context"
	"crm-glonass/api"
	"crm-glonass/config"
	"crm-glonass/data/cache"
	mongox "crm-glonass/data/mongox"
	"fmt"

	"crm-glonass/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	// Logger info
	logger.Info(logging.General, logging.StartUp, "Started server...", map[logging.ExtraKey]interface{}{"Version": conf.Version})

	// Database connection
	database, _ := mongox.Connection(conf, ctx, logger)
	cache.InitRedis(conf, ctx)
	logger.Debug(logging.Swagger, logging.Link, fmt.Sprintf("http://localhost:%d/swagger/index.html", conf.Server.IPort), map[logging.ExtraKey]interface{}{"Version": conf.Version})
	api.InitialServer(conf, database, logger)

}

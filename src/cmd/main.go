package main

import (
	"crm-glonass/src/config"
	mongox "crm-glonass/src/data/db"
	"crm-glonass/src/pkg/logging"
)

var logcod = logging.NewLogger(config.GetConfig())

func main() {

	logcod.Info(logging.General, logging.StartUp, "Started server...", nil)

	conf := config.GetConfig()

	mongox.Connection(conf, logcod)

	select {}
}

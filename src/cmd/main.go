package main

import (
	"crm-glonass/src/config"
	mongox "crm-glonass/src/data/db"
	"crm-glonass/src/pkg/logcod"
)

var logger = logcod.NewLogger(config.GetConfig())

func main() {

	logger.Info(logcod.General, logcod.StartUp, "Started", nil)
	conf := config.GetConfig()

	mongox.Connection(conf)

	select {}
}

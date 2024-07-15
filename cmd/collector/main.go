package main

import (
	"crm-glonass/config"
	"crm-glonass/pkg/logging"
)

func main() {

	logger := logging.NewLogger(config.GetConfig())
	logger.Info(logging.General, logging.StartUp, "Started server...", map[logging.ExtraKey]interface{}{"Version": config.GetConfig().Version})

}

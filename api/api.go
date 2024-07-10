package api

import (
	routers "crm-glonass/api/routes"
	"crm-glonass/config"
	_ "crm-glonass/docs"
	"crm-glonass/middlewares"
	"crm-glonass/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logcod = logging.NewLogger(config.GetConfig())

func InitialServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler), middlewares.LimitByRequest())

	RegisterRouter(r)

	logcod.Info(logging.API, logging.StartUp, "Started API", nil)
	err := r.Run(fmt.Sprintf(":%d", cfg.Server.IPort))
	if err != nil {
		logcod.Fatal(logging.API, logging.StartUp, err.Error(), nil)
	}
}

// @title           CRM COMECORD [DEV]
// @version         0.1.0
// @description     Система управления и мониторинга транспортных средств с системой GLONASS
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alexander Hunter
// @contact.url    http://www.swagger.io/support
// @contact.email  icehuntmen@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5100
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func RegisterRouter(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

package api

import (
	routers "crm-glonass/api/routes"
	"crm-glonass/config"
	"crm-glonass/docs"
	_ "crm-glonass/docs"
	"crm-glonass/middlewares"
	"crm-glonass/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitialServer(cfg *config.Config, database *mongo.Database, logger logging.Logger) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler), middlewares.LimitByRequest())

	RegisterRouter(r, cfg, database)
	RegisterSwagger(r, cfg)

	logger.Info(logging.API, logging.StartUp, "Started API", nil)
	err := r.Run(fmt.Sprintf(":%d", cfg.Server.IPort))
	if err != nil {
		logger.Fatal(logging.API, logging.StartUp, err.Error(), nil)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "COMECORD"
	docs.SwaggerInfo.Description = "Система управление и мониторинга транспортных средст с системой GLONASS"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Server.EPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func RegisterRouter(r *gin.Engine, conf *config.Config, db *mongo.Database) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{

		//vehiclesA := v1.Group("/vehicles")
		vehiclesB := v1.Group("/vehicles", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))

		//routers.Vehicles(vehiclesA, db.Collection("vehicles"))
		routers.Vehicles(vehiclesB, db.Collection("vehicles"))
		health := v1.Group("/health")
		routers.Health(health)
	}

}

package controllers

import (
	"context"
	"crm-glonass/models"
	"crm-glonass/pkg/logging"
	"crm-glonass/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type VehiclesController struct {
	service *services.VehicleService
}

func NewVehiclesController(collection *mongo.Collection, ctx context.Context, logger logging.Logger) *VehiclesController {
	service, ok := services.NewVehicleService(collection, ctx, logger).(*services.VehicleService)
	if !ok {
		// handle the error case where the type assertion fails
		return nil
	}
	return &VehiclesController{
		service: service,
	}
}

// Create Vehicle godoc
// @Summary Create a vehicle
// @Description Create a vehicle
// @Tags Vehicles
// @Accept json
// @produces json
// @Param Request body dto.CreateVehicleRequestDTO true "Create a vehicle model"
// @Success 201 {object} components.BaseHttpResponse{result=dto.DBVehicleDTO} "Created response"
// @Failure 400 {object} components.BaseHttpResponse "Bad request"
// @Router /v1/vehicles/ [post]
// @Security AuthBearer
func (pc *VehiclesController) Create(ctx *gin.Context) {
	var vehicle *models.CreateVehicleRequest

	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPost, err := pc.service.Create(vehicle)

	if err != nil {
		//if strings.Contains(err.Error(), "title already exists") {
		//	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
		//	return
		//}
		//
		//ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

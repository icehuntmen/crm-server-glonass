package controllers

import (
	"context"
	"crm-glonass/api/dto"
	"crm-glonass/config"
	"crm-glonass/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type VehiclesController struct {
	service *services.VehicleService
}

func NewVehiclesController(db *mongo.Database, ctx context.Context, conf *config.Config) *VehiclesController {
	service, ok := services.NewVehicleService(db, conf, ctx, "vehicles").(*services.VehicleService)
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
// @Param Request body dto.CreateVehicleRequest true "Create a vehicle model"
// @Success 201 {object} components.BaseHttpResponse{result=dto.DBVehicleDTO} "Created response"
// @Failure 400 {object} components.BaseHttpResponse "Bad request"
// @Router /v1/vehicles/ [post]
// @Security AuthBearer
func (vc *VehiclesController) Create(ctx *gin.Context) {
	var vehicle *dto.CreateVehicleRequest

	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPost, err := vc.service.Create(vehicle)

	if err != nil {
		// handle error
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

func (pc *VehiclesController) Update(ctx *gin.Context) {

}

func (pc *VehiclesController) Delete(ctx *gin.Context) {

}

func (pc *VehiclesController) GetById(ctx *gin.Context) {

}

func (pc *VehiclesController) GetByFilter(ctx *gin.Context) {

}

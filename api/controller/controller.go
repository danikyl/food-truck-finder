package controller

import (
	"net/http"

	"github.com/danikyl/food-truck-finder/api/domain"
	"github.com/danikyl/food-truck-finder/api/mapper"
	"github.com/danikyl/food-truck-finder/core/usecase"
	"github.com/gin-gonic/gin"
)

type apiController struct {
	findNearestFoodTrucksUseCase usecase.FindNearestFoodTrucks
}

func NewController(findNearestFoodTrucksUseCase usecase.FindNearestFoodTrucks) *apiController {
	return &apiController{findNearestFoodTrucksUseCase: findNearestFoodTrucksUseCase}
}

func (c *apiController) handlefindNearestFoodTrucks(context *gin.Context) {
	var userLocationRequest domain.UserLocationRequest
	context.BindJSON(&userLocationRequest)
	context.IndentedJSON(http.StatusOK, c.findNearestFoodTrucksUseCase.Execute(mapper.ConvertLocationRequestToCoreObject(userLocationRequest)))
}

func (c *apiController) StartServer() {
	router := gin.Default()
	router.POST("/findNearestFoodTrucks", c.handlefindNearestFoodTrucks)
	router.Run("localhost:8080")
}

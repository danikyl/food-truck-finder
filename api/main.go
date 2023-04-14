package main

import (
	"github.com/danikyl/food-truck-finder/api/controller"
	"github.com/danikyl/food-truck-finder/core/usecase"
)

func main() {
	//injecting dependencies
	findNearestFoodTrucksUseCase := usecase.NewNearestFoodTrucksImpl()
	apiController := controller.NewController(findNearestFoodTrucksUseCase)

	apiController.StartServer()
}

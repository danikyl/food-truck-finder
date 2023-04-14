package usecase

import "github.com/danikyl/food-truck-finder/core/domain"

type FindNearestFoodTrucks interface {
	Execute(userLocation domain.UserLocation) []domain.FoodTruck
}
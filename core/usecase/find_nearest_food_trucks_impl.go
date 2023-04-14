package usecase

import "github.com/danikyl/food-truck-finder/core/domain"

type findNearestFoodTrucksImpl struct{}

func (*findNearestFoodTrucksImpl) Execute(userLocation domain.UserLocation) []domain.FoodTruck {
	return []domain.FoodTruck{}
}

func NewNearestFoodTrucksImpl() *findNearestFoodTrucksImpl{
	return &findNearestFoodTrucksImpl{}
}
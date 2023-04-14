package mapper

import (
	apiDomain "github.com/danikyl/food-truck-finder/api/domain"
	coreDomain "github.com/danikyl/food-truck-finder/core/domain"
)

func ConvertLocationRequestToCoreObject(request apiDomain.UserLocationRequest) coreDomain.UserLocation {
	return coreDomain.UserLocation{X: request.X, Y: request.Y, Latitude: request.Latitude, Longitude: request.Longitude}
}

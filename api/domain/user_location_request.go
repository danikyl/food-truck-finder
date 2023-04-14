package domain

type UserLocationRequest struct {
	X         float64 `json:"x_location"`
	Y         float64 `json:"y_location"`
	Latitude  float64 `json:"latitude_location"`
	Longitude float64 `json:"longitude_location"`
}

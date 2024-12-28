package models

import "github.com/go-playground/validator/v10"

// House type schema of my dataset
type House struct {
	Address        string  `json:"address" validate:"required"`
	Neighborhood   string  `json:"neighborhood" validate:"required"`
	Bedrooms       int     `json:"bedrooms" validate:"required,min=1"`
	Bathrooms      int     `json:"bathrooms" validate:"required,min=1"`
	SquareMeters   int     `json:"square_meters" validate:"required,min=1"`
	BuildingAge    int     `json:"building_age" validate:"required,min=1"`
	Garden         string  `json:"garden" validate:"required,oneof=Yes No"`
	Garage         string  `json:"garage" validate:"required,oneof=Yes No"`
	Floors         int     `json:"floors" validate:"required,min=1"`
	PropertyType   string  `json:"property_type" validate:"required"`
	HeatingType    string  `json:"heating_type" validate:"required"`
	Balcony        string  `json:"balcony" validate:"required"`
	InteriorStyle  string  `json:"interior_style" validate:"required"`
	View           string  `json:"view" validate:"required"`
	Materials      string  `json:"materials" validate:"required"`
	BuildingStatus string  `json:"building_status" validate:"required"`
	Price          float64 `json:"price" validate:"required,min=0"`
}

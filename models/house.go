package models

// Using House struct of my csv

type House struct {
	Address        string  `json:"address"`
	Neighborhood   string  `json:"neighborhood"`
	Bedrooms       int     `json:"bedrooms"`
	Bathrooms      int     `json:"bathrooms"`
	SquareMeters   int     `json:"square_meters"`
	BuildingAge    int     `json:"building_age"`
	Garden         bool    `json:"garden"`
	Garage         string  `json:"garage"`
	Floors         int     `json:"floors"`
	PropertyType   string  `json:"property_type"`
	HeatingType    string  `json:"heating_type"`
	Balcony        string  `json:"balcony"`
	InteriorStyle  string  `json:"interior_style"`
	View           string  `json:"view"`
	Materials      string  `json:"materials"`
	BuildingStatus string  `json:"building_status"`
	Price          float64 `json:"price"`
}

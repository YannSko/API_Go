// utils/sanitizer.go
package utils

import (
    "strings"
    "html"
    "api_go/models"
)

// Sanitize sanitizes a string input (e.g., trims, escapes HTML)
func Sanitize(input string) string {
    sanitizedInput := strings.TrimSpace(input)      // Remove leading/trailing spaces
    sanitizedInput = html.EscapeString(sanitizedInput) // Escape HTML special characters
    return sanitizedInput
}

// SanitizeHouse sanitizes all fields of the House struct
func SanitizeHouse(house *models.House) {
    house.Address = Sanitize(house.Address)
    house.Neighborhood = Sanitize(house.Neighborhood)
    house.PropertyType = Sanitize(house.PropertyType)
    house.HeatingType = Sanitize(house.HeatingType)
    house.InteriorStyle = Sanitize(house.InteriorStyle)
    house.View = Sanitize(house.View)
    house.Materials = Sanitize(house.Materials)
}

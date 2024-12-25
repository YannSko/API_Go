package utils

import (
	"api_go/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// ReadCSV reads a CSV file and returns a slice of House objects
func ReadCSV(filePath string) ([]models.House, error) {
	// open csv
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read the CSV file, error path:", err)
		return nil, err
	}
	defer file.Close()

	// iniate a csv reader
	reader := csv.NewReader(file)

	// read all records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV", err)
		return nil, err
	}

	// Create a slice to hold the House objects
	var houses []models.House

	// Loop through the records (skip the header)
	for _, record := range records[1:] {
		// Parse each field into the correct type
		price, _ := strconv.ParseFloat(record[16], 64) // Parse price to float64
		bedrooms, _ := strconv.Atoi(record[2])         // Parse bedrooms to int
		bathrooms, _ := strconv.Atoi(record[3])        // Parse bathrooms to int
		squareMeters, _ := strconv.Atoi(record[4])     // Parse square meters to int
		buildingAge, _ := strconv.Atoi(record[5])      // Parse building age to int
		floors, _ := strconv.Atoi(record[8])           // Parse floors to int

		// Create a House struct and append it to the houses slice
		house := models.House{
			Address:        record[0],
			Neighborhood:   record[1],
			Bedrooms:       bedrooms,
			Bathrooms:      bathrooms,
			SquareMeters:   squareMeters,
			BuildingAge:    buildingAge,
			Garden:         record[6],
			Garage:         record[7],
			Floors:         floors,
			PropertyType:   record[9],
			HeatingType:    record[10],
			Balcony:        record[11],
			InteriorStyle:  record[12],
			View:           record[13],
			Materials:      record[14],
			BuildingStatus: record[15],
			Price:          price,
		}

		houses = append(houses, house)
	}

	return houses, nil
}

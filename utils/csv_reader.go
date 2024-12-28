package utils

import (
    "api_go/models"
    "encoding/csv"
    "context"
    "log"
    "os"
    "strconv"
    "github.com/jackc/pgx/v4"
)

// ReadCSVAndInsertIntoDB lit un fichier CSV et insère chaque enregistrement dans la base de données PostgreSQL
func ReadCSVAndInsertIntoDB(filePath string, db *pgx.Conn) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    // Pour chaque ligne dans le CSV, insérer dans PostgreSQL
    for _, record := range records[1:] { // Ignore le header
        price, _ := strconv.ParseFloat(record[16], 64)
        bedrooms, _ := strconv.Atoi(record[2])
        bathrooms, _ := strconv.Atoi(record[3])
        squareMeters, _ := strconv.Atoi(record[4])
        buildingAge, _ := strconv.Atoi(record[5])
        floors, _ := strconv.Atoi(record[8])

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

        // Insertion dans la base de données PostgreSQL
        _, err = db.Exec(context.Background(), `
            INSERT INTO houses (address, neighborhood, bedrooms, bathrooms, square_meters, building_age, garden, garage, floors, property_type, heating_type, balcony, interior_style, view, materials, building_status, price) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
            house.Address, house.Neighborhood, house.Bedrooms, house.Bathrooms, house.SquareMeters, house.BuildingAge,
            house.Garden, house.Garage, house.Floors, house.PropertyType, house.HeatingType, house.Balcony,
            house.InteriorStyle, house.View, house.Materials, house.BuildingStatus, house.Price)

        if err != nil {
            log.Printf("Error inserting house: %v", err)
            return err
        }
    }
    return nil
}

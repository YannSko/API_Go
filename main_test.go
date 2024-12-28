package main

import (
	"api_go/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	// Créer un Gin router pour les tests
	r := gin.Default()

	// Initialisation des routes
	r.POST("/houses", addHouse)
	r.GET("/houses", getHouses)
	r.PUT("/houses/:id", updateHouse)
	r.DELETE("/houses/:id", deleteHouse)

	// Test POST /houses - Ajouter une maison
	t.Run("POST /houses", func(t *testing.T) {
		newHouse := models.House{
			Address:        "123 Baker Street",
			Neighborhood:   "Marylebone",
			Bedrooms:       4,
			Bathrooms:      2,
			SquareMeters:   200,
			BuildingAge:    50,
			Garden:         "Yes",
			Garage:         "No",
			Floors:         3,
			PropertyType:   "Detached",
			HeatingType:    "Gas Heating",
			Balcony:        "High-level Balcony",
			InteriorStyle:  "Modern",
			View:           "City",
			Materials:      "Wood",
			BuildingStatus: "Renovated",
			Price:          1500000,
		}

		jsonData, _ := json.Marshal(newHouse)
		req := httptest.NewRequest(http.MethodPost, "/houses", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	// Test GET /houses - Récupérer toutes les maisons
	t.Run("GET /houses", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/houses?page=1&pageSize=10", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test PUT /houses/:id - Mettre à jour une maison
	t.Run("PUT /houses/:id", func(t *testing.T) {
		// Simuler une mise à jour
		updatedHouse := models.House{
			Address:        "456 Regent Street",
			Neighborhood:   "Soho",
			Bedrooms:       5,
			Bathrooms:      3,
			SquareMeters:   250,
			BuildingAge:    80,
			Garden:         "Yes",
			Garage:         "Yes",
			Floors:         4,
			PropertyType:   "Semi-Detached",
			HeatingType:    "Electric Heating",
			Balcony:        "Low-level Balcony",
			InteriorStyle:  "Minimalist",
			View:           "Street",
			Materials:      "Marble",
			BuildingStatus: "New",
			Price:          2000000,
		}

		// ID fictif pour la mise à jour
		id := "1"
		jsonData, _ := json.Marshal(updatedHouse)
		req := httptest.NewRequest(http.MethodPut, "/houses/"+id, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test DELETE /houses/:id - Supprimer une maison
	t.Run("DELETE /houses/:id", func(t *testing.T) {
		// ID fictif pour la suppression
		id := "1"
		req := httptest.NewRequest(http.MethodDelete, "/houses/"+id, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})
}


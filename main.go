package main

import (
	"api_go/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.GET("/houses", func(c *gin.Context) {
		// Load csv file house
		houses, err := utils.ReadCSV("data/london_houses.csv")
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Unable to read the CSV file",
			})
			return
		}
		c.JSON(200, houses)
	})
	r.Run(":8080")
	log.Println("Server running on port 8080")

}

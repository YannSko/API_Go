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
        houses, err := utils.ReadCSV("data/london_houses.csv")
        if err != nil {
            c.JSON(500, gin.H{
                "error": "Unable to read the CSV file",
            })
            return
        }
        c.JSON(200, houses)
    })

    r.POST("/houses", func(c *gin.Context) {
        var newHouse utils.House
        if err := c.ShouldBindJSON(&newHouse); err != nil {
            c.JSON(400, gin.H{
                "error": "Invalid input",
            })
            return
        }
        // Here you would typically add the new house to your data store
        c.JSON(201, newHouse)
    })

    r.PUT("/houses/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updatedHouse utils.House
        if err := c.ShouldBindJSON(&updatedHouse); err != nil {
            c.JSON(400, gin.H{
                "error": "Invalid input",
            })
            return
        }
        // Here you would typically update the house with the given id in your data store
        c.JSON(200, gin.H{
            "id":      id,
            "updated": updatedHouse,
        })
    })

    r.DELETE("/houses/:id", func(c *gin.Context) {
        id := c.Param("id")
        // Here you would typically delete the house with the given id from your data store
        c.JSON(200, gin.H{
            "id":      id,
            "deleted": true,
        })
    })

    r.Run(":8080")
    log.Println("Server running on port 8080")
}
package main

import (
    "api_go/models"
    "api_go/utils"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/juju/ratelimit"
)

func main() {
    r := gin.Default()
    limiter := ratelimit.NewBucketWithRate(100, 100)

    r.Use(func(c *gin.Context) {
        if limiter.TakeAvailable(1) == 0 {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    })
    
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

        // Get pagination parameters
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

        // Calculate start and end indices
        start := (page - 1) * pageSize
        end := start + pageSize

        // Ensure indices are within bounds
        if start > len(houses) {
            start = len(houses)
        }
        if end > len(houses) {
            end = len(houses)
        }

        // Return paginated results
        c.JSON(200, houses[start:end])
    })

    r.POST("/houses", func(c *gin.Context) {
        var newHouse models.House
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
        var updatedHouse models.House
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
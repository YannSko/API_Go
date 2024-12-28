package main

import (
    "api_go/models"
    "api_go/utils"
    "log"
    "net/http"
    "strconv"
    "os"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/juju/ratelimit"
    "github.com/jackc/pgx/v4"
    "context"
    "github.com/go-playground/validator/v10"
)

var db *pgx.Conn
var validate *validator.Validate

func main() {
    // Connexion à la base de données
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", dbUser, dbPassword, dbHost, dbName)

    var err error
    db, err = pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer db.Close(context.Background()) // Fermer la connexion à la fin

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

    validate = validator.New()

    // Test de connexion à l'API (Pas besoin de JWT pour cette route)
    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "API running",
        })
    })

    // Route de login pour récupérer un token JWT (Pas besoin de JWT pour cette route)
    r.POST("/login", func(c *gin.Context) {
        var loginDetails struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&loginDetails); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }

        // Exemple d'utilisateur (à remplacer par une vérification dans ta base de données)
        if loginDetails.Username == "admin" && loginDetails.Password == "password" {
            token, err := utils.GenerateJWT("12345") // Remplacer par un ID d'utilisateur réel
            if err != nil {
                c.JSON(500, gin.H{"error": "Could not generate token"})
                return
            }
            c.JSON(200, gin.H{"token": token})
        } else {
            c.JSON(401, gin.H{"error": "Invalid credentials"})
        }
    })

    // Appliquer le middleware JWT uniquement aux routes protégées
    authorized := r.Group("/")
    authorized.Use(func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        // Valider le token JWT
        _, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    })

    // Routes protégées : requièrent un token JWT valide
    authorized.GET("/houses", func(c *gin.Context) {
        // Paramètres de pagination
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

        start := (page - 1) * pageSize
        end := start + pageSize

        // Récupérer les maisons depuis PostgreSQL
        rows, err := db.Query(context.Background(), `SELECT id, address, neighborhood, bedrooms, bathrooms, square_meters, building_age, garden, garage, floors, property_type, heating_type, balcony, interior_style, view, materials, building_status, price FROM houses LIMIT $1 OFFSET $2`, pageSize, start)
        if err != nil {
            c.JSON(500, gin.H{
                "error": "Unable to fetch houses from database",
            })
            return
        }
        defer rows.Close()

        var houses []models.House
        for rows.Next() {
            var house models.House
            err := rows.Scan(&house.ID, &house.Address, &house.Neighborhood, &house.Bedrooms, &house.Bathrooms, &house.SquareMeters, &house.BuildingAge, &house.Garden, &house.Garage, &house.Floors, &house.PropertyType, &house.HeatingType, &house.Balcony, &house.InteriorStyle, &house.View, &house.Materials, &house.BuildingStatus, &house.Price)
            if err != nil {
                c.JSON(500, gin.H{
                    "error": "Error scanning house data",
                })
                return
            }
            houses = append(houses, house)
        }

        if err := rows.Err(); err != nil {
            c.JSON(500, gin.H{
                "error": "Error fetching houses from database",
            })
            return
        }

        // Retourner les maisons en format JSON
        c.JSON(200, houses)
    })

    // Ajouter une maison dans la base de données PostgreSQL
    authorized.POST("/houses", func(c *gin.Context) {
        var newHouse models.House
        if err := c.ShouldBindJSON(&newHouse); err != nil {
            c.JSON(400, gin.H{
                "error": "Invalid input",
            })
            return
        }

        // Sanitize input before validation
        utils.SanitizeHouse(&newHouse)

        // Validation des données
        if err := validate.Struct(newHouse); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }

        // Ajouter la maison dans la base de données PostgreSQL
        _, err := db.Exec(context.Background(), `
            INSERT INTO houses (address, neighborhood, bedrooms, bathrooms, square_meters, building_age, garden, garage, floors, property_type, heating_type, balcony, interior_style, view, materials, building_status, price) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
            newHouse.Address, newHouse.Neighborhood, newHouse.Bedrooms, newHouse.Bathrooms, newHouse.SquareMeters, newHouse.BuildingAge,
            newHouse.Garden, newHouse.Garage, newHouse.Floors, newHouse.PropertyType, newHouse.HeatingType, newHouse.Balcony,
            newHouse.InteriorStyle, newHouse.View, newHouse.Materials, newHouse.BuildingStatus, newHouse.Price)

        if err != nil {
            c.JSON(500, gin.H{"error": "Unable to add house to database"})
            return
        }

        c.JSON(201, newHouse)
    })

    // Mettre à jour une maison dans la base de données
    authorized.PUT("/houses/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updatedHouse models.House
        if err := c.ShouldBindJSON(&updatedHouse); err != nil {
            c.JSON(400, gin.H{
                "error": "Invalid input",
            })
            return
        }
        
        // Sanitize input before validation
        utils.SanitizeHouse(&updatedHouse)

        // Mise à jour dans la base de données
        _, err := db.Exec(context.Background(), `UPDATE houses SET address=$1, neighborhood=$2, bedrooms=$3, bathrooms=$4, square_meters=$5, building_age=$6, garden=$7, garage=$8, floors=$9, property_type=$10, heating_type=$11, balcony=$12, interior_style=$13, view=$14, materials=$15, building_status=$16, price=$17 WHERE id=$18`,
            updatedHouse.Address, updatedHouse.Neighborhood, updatedHouse.Bedrooms, updatedHouse.Bathrooms, updatedHouse.SquareMeters, updatedHouse.BuildingAge,
            updatedHouse.Garden, updatedHouse.Garage, updatedHouse.Floors, updatedHouse.PropertyType, updatedHouse.HeatingType, updatedHouse.Balcony,
            updatedHouse.InteriorStyle, updatedHouse.View, updatedHouse.Materials, updatedHouse.BuildingStatus, updatedHouse.Price, id)

        if err != nil {
            c.JSON(500, gin.H{"error": "Unable to update house"})
            return
        }

        c.JSON(200, gin.H{
            "id":      id,
            "updated": updatedHouse,
        })
    })

    // Supprimer une maison de la base de données
    authorized.DELETE("/houses/:id", func(c *gin.Context) {
        id := c.Param("id")

        // Supprimer la maison de la base de données
        _, err := db.Exec(context.Background(), "DELETE FROM houses WHERE id=$1", id)
        if err != nil {
            c.JSON(500, gin.H{"error": "Unable to delete house"})
            return
        }

        c.JSON(200, gin.H{
            "id":      id,
            "deleted": true,
        })
    })

    // Démarrer le serveur sur le port 8080
    r.Run(":8080")
    log.Println("Server running on port 8080")
}

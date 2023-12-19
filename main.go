package main

import (
	// "errors"
	"os"
	// "fmt"
	"log"
	// "net/http"
	// "regexp"
	// "strconv"
	// "time"
	// "context"
	// "math/rand"
	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "gopkg.in/gomail.v2"
	"github.com/jamyMarkos/backend/routes"
	"github.com/jamyMarkos/backend/models"
	"github.com/jamyMarkos/backend/middleware"
)


func init() {
    // Load environmental variables from .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

// Driver Code...
func main() {

	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	models.ConnectToMongoDB()

	// Set up routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is starting on port %s...\n", port)

	r.Run(":" + port)
}

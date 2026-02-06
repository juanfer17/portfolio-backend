package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/repository"
	"portfolio-backend/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "portfolio" // Default database name
	}

	// Connect to MongoDB
	client, err := database.ConnectDB(mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// --- Dependency Injection ---
	db := client.Database(dbName)
	repo := repository.NewMongoRepository(db)
	contactSvc := service.NewContactService(repo)
	handler := handlers.NewPortfolioHandler(repo, contactSvc)

	// --- Router Setup ---
	r := gin.Default()

	// 1. Global Middleware: CORS
	// Applied to ALL routes, including health check and public endpoints
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://portfolio-frontend-green-tau.vercel.app", // Production Frontend
			"http://localhost:3000",                           // Local Development
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "X-API-KEY"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 2. Public Routes (No Auth Required)

	// Health Check (For UptimeRobot / Render)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// Read-only endpoints for the Portfolio Frontend
	r.GET("/tech", handler.GetTechnologies)
	r.GET("/experience", handler.GetExperience)
	r.POST("/contact", handler.SendContact)

	// 3. Protected Routes (Require X-API-KEY)
	// Used only by the Admin/Developer to manage content
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware())
	{
		// Technology CRUD (Write)
		admin.POST("/tech", handler.CreateTechnology)
		admin.PUT("/tech/:id", handler.UpdateTechnology)
		admin.DELETE("/tech/:id", handler.DeleteTechnology)

		// Experience CRUD (Write)
		admin.POST("/experience", handler.CreateExperience)
		admin.PUT("/experience/:id", handler.UpdateExperience)
		admin.DELETE("/experience/:id", handler.DeleteExperience)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

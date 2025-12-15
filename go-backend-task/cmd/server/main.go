package main

import (
	"context"
	"fmt"
	"os"

	"go-backend-task/internal/handler"
	"go-backend-task/internal/logger"
	"go-backend-task/internal/middleware"
	"go-backend-task/internal/repository"
	"go-backend-task/internal/routes"
	"go-backend-task/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" // Import this
)

func main() {
	// 1. Initialize Logger
	logger.InitLogger()

	// 2. Load .env file
	// We allow this to fail silently in case we are in a production environment 
	// without a .env file, but for local dev, it loads variables.
	// Try loading from the current directory, or explicit path if needed
if err := godotenv.Load(); err != nil {
    // If that fails, try loading explicitly from the root (common fix for VS Code terminals)
    if err := godotenv.Load("../../.env"); err != nil {
         logger.Log.Info("No .env file found")
    }
}

	// 3. Get DB Source from Environment Variable
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		logger.Log.Fatal("DB_SOURCE environment variable is not set")
	}

	// 4. Database Connection
	dbPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		logger.Log.Fatal("Unable to connect to database: " + err.Error())
	}
	defer dbPool.Close()

	// 5. Initialize Dependency Injection
	store := repository.NewStore(dbPool)
	userService := service.NewUserService(store)
	userHandler := handler.NewUserHandler(userService)

	// 6. Setup Fiber
	app := fiber.New()
	middleware.SetupMiddleware(app)
	routes.SetupRoutes(app, userHandler)

	// 7. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	fmt.Println("Server running on port " + port)
	logger.Log.Fatal(app.Listen(":" + port).Error())
}
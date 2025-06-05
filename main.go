package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bhuwan-darai/crud/config"
	"github.com/bhuwan-darai/crud/graph"
	"github.com/bhuwan-darai/crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println(DATABASE_URL)
	fmt.Printf("Server started on port %s\n", port)

	// initialize fiber
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Vite dev server origin
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	// Connect to DB
	db := config.ConnectDB()
	defer db.Prisma.Disconnect()

	// Create a GraphQL Resolver and pass DB client
	resolver := &graph.Resolver{
		DB: db,
	}

	// Set up GraphQL and Playground routes
	routes.SetupRoutes(app, resolver)

	// Home route
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println(("Server started on port 8000"))

	// Start server
	app.Listen(":" + port)
}

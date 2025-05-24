package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bhuwan-darai/crud/config"
	"github.com/bhuwan-darai/crud/graph"
	"github.com/bhuwan-darai/crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	fmt.Println(DATABASE_URL)

	app := fiber.New()

	db := config.ConnectDB()
	defer db.Prisma.Disconnect()

	// Create a GraphQL Resolver and pass DB client
	resolver := &graph.Resolver{
		DB: db,
	}

	// Set up GraphQL and Playground routes
	routes.SetupRoutes(app, resolver)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Bhuwan-Darai/goCrud/config"
	"github.com/Bhuwan-Darai/goCrud/graph"
	"github.com/Bhuwan-Darai/goCrud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

//go:generate go run github.com/steebchen/prisma-client-go generate

func main() {
	log.Println("PORT:", os.Getenv("PORT"))
	log.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file loaded (probably production):", err)
		}
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")
	fmt.Println(PORT)

	if PORT == "" {
		PORT = defaultPort
	}

	fmt.Println(DATABASE_URL)
	// fmt.Printf("Server started on port %s\n", PORT)

	// initialize fiber
	app := fiber.New()

	// CORS middleware
	env := os.Getenv("ENV")
	var origins string

	if env == "production" {
		origins = "https://go-server-todo-frontend.vercel.app"
	} else {
		// default to development
		origins = "http://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers",
	}))

	// actual origin used
	fmt.Println("Origins:", origins)

	// Debug: Check if Prisma engine files exist
	checkPrismaEngine()

	// Connect to DB
	db := config.ConnectDB()
	defer db.Prisma.Disconnect()

	log.Println("✅ Connected to Supabase PostgreSQL")

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

	// Add health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Log server startup
	// log.Printf("Server starting on port %s", PORT)

	log.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(app.Listen(":" + PORT))
}

func checkPrismaEngine() {
	paths := []string{
		"./prisma/client/engine/",
		"./prisma/db/prisma-client/",
		"./prisma/",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("✅ Found Prisma directory: %s\n", path)

			// List files in the directory
			err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					fmt.Printf("  - %s\n", filePath)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("❌ Error listing files in %s: %v\n", path, err)
			}
		} else {
			fmt.Printf("❌ Prisma directory not found: %s\n", path)
		}
	}
}

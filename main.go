package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Bhuwan-Darai/goCrud/config"
	"github.com/Bhuwan-Darai/goCrud/graph"
	"github.com/Bhuwan-Darai/goCrud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

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
	origins := "http://localhost:3000, https://go-server-todo-frontend.vercel.app/"
	app.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
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

	// Log server startup
	// log.Printf("Server starting on port %s", PORT)

	log.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(app.Listen(":" + PORT))
}

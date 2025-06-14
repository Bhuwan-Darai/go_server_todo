package config

import (
	"log"
	"os"
	"strings"

	db "github.com/Bhuwan-Darai/goCrud/prisma/db/prisma-client"
)

func ConnectDB() *db.PrismaClient {
	database := db.NewClient()
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment")
	}

	// Replace pgbouncer=true with sslmode=require for direct connection
	if os.Getenv("PRISMA_DIRECT_CONNECTION") == "true" {
		databaseURL = strings.Replace(databaseURL, "pgbouncer=true", "sslmode=require", 1)
	}

	log.Printf("📦 Connecting to DB: %s", databaseURL)

	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}

	log.Println("✅ Successfully connected to database")
	return database
}

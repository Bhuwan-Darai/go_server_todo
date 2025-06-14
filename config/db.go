package config

import (
	"log"
	"os"
	"path/filepath"

	db "github.com/Bhuwan-Darai/goCrud/prisma/db/prisma-client"
)

func ConnectDB() *db.PrismaClient {
	database := db.NewClient()
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment")
	}

	log.Printf("📦 Connecting to DB: %s", databaseURL)

	// Verify Prisma client files exist
	prismaDir := "./prisma"
	clientDir := filepath.Join(prismaDir, "db/prisma-client")
	queryEngine := filepath.Join(prismaDir, "query-engine")

	if _, err := os.Stat(clientDir); os.IsNotExist(err) {
		log.Fatalf("❌ Prisma client directory not found at: %s", clientDir)
	}

	if _, err := os.Stat(queryEngine); os.IsNotExist(err) {
		log.Fatalf("❌ Prisma query engine not found at: %s", queryEngine)
	}

	log.Printf("✅ Found Prisma client at: %s", clientDir)
	log.Printf("✅ Found query engine at: %s", queryEngine)

	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}

	log.Println("✅ Successfully connected to database")
	return database
}

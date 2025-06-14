package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

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
	queryEngine := filepath.Join(clientDir, "query-engine")

	// Check client directory
	if _, err := os.Stat(clientDir); os.IsNotExist(err) {
		log.Fatalf("❌ Prisma client directory not found at: %s", clientDir)
	}

	// Check query engine
	if _, err := os.Stat(queryEngine); os.IsNotExist(err) {
		log.Fatalf("❌ Query engine not found at: %s", queryEngine)
	}

	// Ensure query engine is executable
	if err := os.Chmod(queryEngine, 0755); err != nil {
		log.Printf("⚠️ Warning: Failed to set query engine permissions: %v", err)
	}

	log.Printf("✅ Found Prisma client at: %s", clientDir)
	log.Printf("✅ Found query engine at: %s", queryEngine)

	// Try to connect with retries
	var err error
	for i := 0; i < 5; i++ {
		if err = database.Prisma.Connect(); err == nil {
			log.Println("✅ Successfully connected to database")
			return database
		}
		log.Printf("⚠️ Connection attempt %d failed: %v", i+1, err)
		time.Sleep(time.Second * 2)
	}

	log.Fatalf("❌ Failed to connect to database after 5 attempts: %v", err)
	return database
}

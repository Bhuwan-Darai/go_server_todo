package config

import (
	"log"
	"os"
	"os/exec"
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
	queryEngine := filepath.Join(clientDir, "query-engine")

	// Check client directory
	if _, err := os.Stat(clientDir); os.IsNotExist(err) {
		log.Fatalf("❌ Prisma client directory not found at: %s", clientDir)
	}

	// Check query engine
	if _, err := os.Stat(queryEngine); os.IsNotExist(err) {
		// Try to find any query engine file
		files, _ := filepath.Glob(filepath.Join(prismaDir, "**/query-engine*"))
		if len(files) > 0 {
			log.Printf("Found query engine files: %v", files)
			// Try to use the first found query engine
			if err := os.Symlink(files[0], queryEngine); err != nil {
				log.Printf("Failed to create symlink: %v", err)
			}
		} else {
			// Try to download the query engine
			log.Printf("Attempting to download query engine...")
			cmd := exec.Command("curl", "-L",
				"https://binaries.prisma.sh/all_commits/5e91ac6b6a6fd5269b456d6063faed3d59a1700c/linux-musl/query-engine.gz",
				"-o", queryEngine+".gz")
			if err := cmd.Run(); err != nil {
				log.Printf("Failed to download query engine: %v", err)
			} else {
				cmd = exec.Command("gunzip", queryEngine+".gz")
				if err := cmd.Run(); err != nil {
					log.Printf("Failed to extract query engine: %v", err)
				} else {
					os.Chmod(queryEngine, 0755)
				}
			}
		}
	}

	// Check query engine permissions
	if info, err := os.Stat(queryEngine); err == nil {
		log.Printf("Query engine permissions: %v", info.Mode())
		if info.Mode()&0111 == 0 {
			log.Printf("Query engine is not executable, attempting to fix...")
			if err := os.Chmod(queryEngine, 0755); err != nil {
				log.Printf("Failed to make query engine executable: %v", err)
			}
		}
	}

	log.Printf("✅ Found Prisma client at: %s", clientDir)
	log.Printf("✅ Found query engine at: %s", queryEngine)

	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}

	log.Println("✅ Successfully connected to database")
	return database
}

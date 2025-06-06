package config

import (
	"log"

	// prisma "github.com/bhuwan-darai/crud/prisma/db"
	"github.com/bhuwan-darai/crud/prisma/db"
)

// DB holds the Prisma client

func ConnectDB() *db.PrismaClient {
	database := db.NewClient()
	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("Error connecting database :%v", err)
	}
	return database
}

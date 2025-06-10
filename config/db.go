package config

import (
	"log"

	// prisma "github.com/bhuwan-darai/crud/prisma/db"
	// "github.com/bhuwan-darai/crud/prisma/db"
	db "github.com/Bhuwan-Darai/goCrud/prisma/db/prisma-client"
)

// DB holds the Prisma clientss

func ConnectDB() *db.PrismaClient {
	database := db.NewClient()
	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("Error connecting database :%v", err)
	}
	return database
}

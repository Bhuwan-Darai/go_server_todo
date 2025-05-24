package config

import (
	"log"

	prisma "github.com/bhuwan-darai/crud/prisma/db"
)

// DB holds the Prisma client

func ConnectDB() *prisma.PrismaClient {
	db := prisma.NewClient()
	if err := db.Prisma.Connect(); err != nil {
		log.Fatalf("Error connecting database :%v", err)
	}
	return db
}

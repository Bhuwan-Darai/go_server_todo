# 🧬 GraphQL Fiber Postgres Boilerplate

A full-featured GraphQL server built with **Go Fiber**, **gqlgen**, **Prisma**, and **PostgreSQL**.

## 🧱 Stack

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/) – Fast HTTP framework
- [gqlgen](https://gqlgen.com/) – GraphQL server generator
- [Prisma](https://www.prisma.io/) – Type-safe database ORM
- [PostgreSQL](https://www.postgresql.org/) – Relational database

---

## 📁 Project Structure

graphql-fiber-postgres/
├── generated/
│ └── prisma-client/ # Generated Prisma Go client
├── graph/
│ ├── generated/ # gqlgen-generated GraphQL code
│ ├── model/ # GraphQL model types
│ ├── schema.graphqls # GraphQL schema definition
│ └── schema.resolvers.go # Custom resolvers (interacting with Prisma)
├── prisma/
│ ├── schema.prisma # Prisma schema (models, db config)
│ └── migrations/ # Prisma migration history
├── routes/
│ └── routes.go # Fiber route definitions (REST or middleware)
├── .env # Environment variables (e.g., DATABASE_URL)
├── go.mod # Go module file
├── go.sum # Go dependencies checksums
├── main.go # Main server entry point
└── README.md # This file

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/graphql-fiber-postgres.git
cd graphql-fiber-postgres

```

### 2. Create a .env file

DATABASE_URL="postgresql://user:password@localhost:5432/dbname?sslmode=disable"

### 3. Install Dependencies

go mod tidy

### 4. Initialize Prisma

npx prisma generate
npx prisma migrate dev --name init

### 5. Generate GraphQL Code

go run github.com/99designs/gqlgen generate

### 6. Run the server

go run main.go

### Note :

- Server runs on: http://localhost:4000

- GraphQL Playground: http://localhost:4000/graphql

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

![project structure](public/go project structure.png)

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

- Server runs on: http://localhost:3000

- GraphQL Playground: http://localhost:3000/graphql

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

![project structure](public/project_structure.png)

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

## Test the server running status in browser

- localhost:8000/health || localhost:8000/hello

## 🚀 Deployment

### Docker Deployment

1. Build the Docker image:

```bash
docker build -t go-todo-server .
```

2. Run the container:

```bash
docker run -p 8000:8000 --env-file .env go-todo-server
```

### Data Flow Architecture

The application follows a layered architecture:

1. **API Layer (GraphQL)**

   - Entry point for all client requests
   - Handles GraphQL queries and mutations
   - Located in `graph/` directory

2. **Service Layer**

   - Contains business logic
   - Processes data between API and database
   - Implements resolvers for GraphQL operations

3. **Database Layer (Prisma)**
   - Manages database operations
   - Provides type-safe database access
   - Handles migrations and schema changes

### Environment Variables

Required environment variables for deployment:

```
DATABASE_URL="postgresql://user:password@localhost:5432/dbname?sslmode=disable"
PORT=8000
ENV=production
```

### Production Considerations

1. **Database**

   - Use a managed PostgreSQL service in production
   - Ensure proper backup and monitoring
   - Configure connection pooling

2. **Security**

   - Enable SSL/TLS for database connections
   - Implement proper authentication
   - Use secure environment variables

3. **Performance**

   - Configure appropriate connection pools
   - Enable query caching where applicable
   - Monitor server resources

4. **Monitoring**
   - Set up logging and monitoring
   - Track database performance
   - Monitor API response times

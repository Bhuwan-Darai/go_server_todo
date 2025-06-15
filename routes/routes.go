package routes

import (
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Bhuwan-Darai/goCrud/graph"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/vektah/gqlparser/v2/ast"
)

// SetupRoutes configures Fiber routes for GraphQL and Playground
func SetupRoutes(app *fiber.App, resolver *graph.Resolver) {
	// GraphQL handler
	graphqlHandler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Configure transports
	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.GET{})
	graphqlHandler.AddTransport(transport.Options{})
	graphqlHandler.AddTransport(transport.MultipartForm{})

	// Enable extensions (similar to NewDefaultServer)
	graphqlHandler.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	graphqlHandler.Use(extension.Introspection{})
	graphqlHandler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// GraphQL endpoint - handle all methods
	app.Use("/graphql", func(c *fiber.Ctx) error {
		// Handle OPTIONS request
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return adaptor.HTTPHandler(graphqlHandler)(c)
	})

	// GraphQL Playground (for development)
	app.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("GraphQL Playground", "/graphql")))
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/gqlgen"
	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

const (
	API_URL = "${API_HOST}:${API_PORT}"
	DB_URI  = "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable"
)

func main() {
	ctx := context.Background()

	// initialize database
	db, err := pgx.Open(ctx, os.ExpandEnv(DB_URI))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize repository
	repoSvc := pgx.NewRepository(db)

	// initialize router
	router := gin.Default()
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler(repoSvc))

	// initialize server
	server := &http.Server{
		Addr:           os.ExpandEnv(API_URL),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Defining the Graphql handler
func graphqlHandler(repoService pgx.Repository) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.New(gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers: &gqlgen.Resolver{
			Repository: repoService,
		},
	}))

	// Server setup:
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

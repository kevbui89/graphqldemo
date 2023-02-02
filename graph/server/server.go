package main

import (
	"log"
	"net/http"
	"os"

	"com.example/graphql/graph/domain"
	"com.example/graphql/graph/graphql"
	graph "com.example/graphql/graph/graphql"
	custommiddleware "com.example/graphql/graph/middleware"
	"com.example/graphql/graph/postgres"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "meetup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UserRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(custommiddleware.AuthMiddleware(userRepo))

	d := domain.NewDomain(userRepo, postgres.MeetupsRepo{DB: DB})

	// The repositories need to be injected in the Resolver
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Domain: d,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	router.Handle("/query", graphql.DataLoaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

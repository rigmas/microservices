package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rigmas/microservices/gateway/graph"
	"github.com/rigmas/microservices/gateway/graph/generated"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r := chi.NewRouter() //Initialize chi router

	r.Use(cors.AllowAll().Handler) //allow cors

	r.Handle("/", playground.Handler("GraphQL playground", "query"))
	r.Handle("/query", srv)

	log.Printf("Server is running on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

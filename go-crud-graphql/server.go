package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sagarshukla785/go-crud-graphql/graph"
	"github.com/sagarshukla785/go-crud-graphql/database"
	"github.com/sagarshukla785/go-crud-graphql/graph/model"
)

const defaultPort = "8080"

func main() {
	db := database.ConnectToMyDB()
    db.AutoMigrate(&model.JobListing{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

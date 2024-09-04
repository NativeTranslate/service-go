package main

import (
	"auth"
	"config"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"graph"
	"log"
	database "mysql"
	"net/http"
)

func main() {
	conf, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	database.InitDB(conf)
	defer database.CloseDB()

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := conf.Server.Port
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

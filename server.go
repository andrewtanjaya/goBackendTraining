package main

import (
	"github.com/andrewtanjaya21/test_go/postgres"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andrewtanjaya21/test_go/graph"
	"github.com/andrewtanjaya21/test_go/graph/generated"
)

const defaultPort = "8080"

func main() {

	//pgDB := pg.Connect(&pg.Options{
	//	Addr:     ":5432",
	//	User:     "postgres",
	//	Password: "root",
	//	Database: "resto",
	//})
	pgDB := postgres.New(
		&pg.Options{
				Addr:     ":5432",
				User:     "postgres",
				Password: "root",
			Database: "resto",
			})
	pgDB.AddQueryHook(postgres.DBLogger{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	defer pgDB.Close()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: pgDB}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

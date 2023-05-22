package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nathanbahiadev/go-graphql/graph"
	"github.com/nathanbahiadev/go-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	fmt.Println("Iniciando aplicação...")

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS categories(id string, name string, description string, primary key(id));"); err != nil {
		log.Fatalf("failed to create table categories: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS courses(id string, name string, description string, category_id string, primary key(id));"); err != nil {
		log.Fatalf("failed to create table courses: %v", err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

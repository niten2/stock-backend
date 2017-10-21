package main

import (
  "fmt"
  "os"
	"net/http"
  "github.com/joho/godotenv"

  "go-todo-api/db"
	"go-todo-api/routers"

  // "go-todo-api/graphql"
	// "github.com/graphql-go/handler"
)

func init() {
  _ = godotenv.Load()
	db.Connect()
}

func main() {
  fmt.Println("server run")

  // mux := http.NewServeMux()

	// h := handler.New(&handler.Config{
		// Schema: &graphql.Schema,
		// Pretty: true,
		// GraphiQL: true,
	// })


  // mux.HandleFunc("/v1", h)

	http.ListenAndServe(os.Getenv("PORT"), routers.Handlers())
}

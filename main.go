package main

import (
  "fmt"
  "os"
	"net/http"
  "github.com/joho/godotenv"

  "go-todo-api/db"
	"go-todo-api/routers"
)

func init() {
  _ = godotenv.Load()
	db.Connect()
}

func main() {
  _ = godotenv.Load()
  fmt.Println("server run")
	http.ListenAndServe(os.Getenv("PORT"), routers.Handlers())
}

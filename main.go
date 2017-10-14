package main

import (
  "fmt"
  "os"
	"net/http"
  "github.com/joho/godotenv"
	"go-todo-api/routers"
)

func main() {
  _ = godotenv.Load()
	http.ListenAndServe(os.Getenv("PORT"), routers.Handlers())
  fmt.Println("server run")
}

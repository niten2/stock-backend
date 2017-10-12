package main

import (
  "fmt"
  "os"
	"encoding/json"
	"net/http"
  "github.com/go-martini/martini"
  "github.com/joho/godotenv"
  "github.com/codegangsta/martini-contrib/render"
  // _ "github.com/joho/godotenv/autoload"
  // "gopkg.in/mgo.v2"

	"go-todo-api/db"
	"go-todo-api/routers"
	"go-todo-api/graphql"
)

func init() {
  _ = godotenv.Load()
	db.Connect()
}

func main() {
  port := os.Getenv("PORT")
	if len(port) == 0 { port = ":3000" }

	m := martini.Classic()

	m.Map(db.Db)
  m.Use(render.Renderer())
  m.Get("/", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"service": "go-todo", "version": "/v1"})
  })
	m.Get("/products", routers.ProductsHandler)

  http.Handle("/", m)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Query().Get("query"))
		result := graphql.ExecuteQuery(r.URL.Query().Get("query"))
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(port, nil)
}

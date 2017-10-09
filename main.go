package main

import (
  "fmt"
  "os"
	"encoding/json"
	"net/http"
  "github.com/go-martini/martini"
  "github.com/joho/godotenv"
  "github.com/codegangsta/martini-contrib/render"
  "gopkg.in/mgo.v2"

	"go-todo-api/routers"
	"go-todo-api/graphql"
)

func main() {
  initGodotenv()

	mongoSession, err := mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
    fmt.Println("Error mongoSession")
		panic(err)
	}

	db := mongoSession.DB(os.Getenv("DB_NAME"))

	m := martini.Classic()

	m.Map(db)
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

	http.ListenAndServe(os.Getenv("PORT"), nil)
}


func initGodotenv () {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
}

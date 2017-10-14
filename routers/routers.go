package routers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "io/ioutil"
  // "bytes"
  "go-todo-api/graphql"
)

// type test_struct struct {


// }


func Handlers() http.Handler {
  r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/v1", graphqlHandler)

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    AllowCredentials: true,
    // Debug: true,
  })

  handler := c.Handler(r)

  return handler
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
  res, _ := json.Marshal(map[string]string{"service": "go-todo", "version": "v1", "graphql": "/v1"})
  w.Write(res)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  // query := r.URL.Query().Get("query")

  query, _ := ioutil.ReadAll(r.Body)
  var dat map[string]string

  if err := json.Unmarshal(query, &dat); err != nil { panic(err) }

  fmt.Println(dat["query"])

  result := graphql.ExecuteQuery(dat["query"])
  json.NewEncoder(w).Encode(result)
}

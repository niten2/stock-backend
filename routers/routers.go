package routers

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "go-todo-api/graphql"
)

func Handlers() *mux.Router {
  r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/v1", graphqlHandler)
  return r
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
  res, _ := json.Marshal(map[string]string{"service": "go-todo", "version": "v1", "graphql": "/v1"})
  w.Write(res)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  result := graphql.ExecuteQuery(r.URL.Query().Get("query"))
  json.NewEncoder(w).Encode(result)
}

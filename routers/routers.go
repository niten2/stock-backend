package routers

import (
  // "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

  "go-todo-api/schema"
  // "go-todo-api/utils"
  // "go-todo-api/utils"
)

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
  res, _ := json.Marshal(map[string]string{
    "service": "go-todo",
    "version": "v1",
    "graphql": "/v1",
  })

  w.Write(res)
}

// func graphqlHandler() http.Handler {
//   h := handler.New(&handler.Config{
//     Schema: &graphql.Schema,
//     Pretty: true,
//   })

//   // return utils.LoggerMiddleware(h)
//   return h
// }


func graphqlHandler(w http.ResponseWriter, r *http.Request) {

  opts := handler.NewRequestOptions(r)

  log.Println("Query", opts.Query)
  log.Println("Variables", opts.Variables)

  params := graphql.Params{
    Schema:         schema.Schema,
    RequestString:  opts.Query,
    VariableValues: opts.Variables,
    OperationName:  opts.OperationName,
    Context:        r.Context(),
  }

  result := graphql.Do(params)

  log.Println("Data", result.Data)
  log.Println("Errors", result.Errors)

  w.Header().Add("Content-Type", "application/json; charset=utf-8")

  w.WriteHeader(http.StatusOK)

  buff, _ := json.MarshalIndent(result, "", "\t")

  w.Write(buff)
}


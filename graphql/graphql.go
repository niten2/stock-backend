package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"

	"go-todo-api/graphql/query"
	"go-todo-api/graphql/mutation"
)

func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: query.Query,
	Mutation: mutation.Mutation,
})

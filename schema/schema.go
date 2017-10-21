package schema

import (
	// "fmt"
	"github.com/graphql-go/graphql"

	"go-todo-api/schema/query"
	"go-todo-api/schema/mutation"
)

// func ExecuteQuery(query string) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema: Schema,
// 		RequestString: query,
// 	})

// 	if len(result.Errors) > 0 {
// 		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
// 	}

// 	return result
// }

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: query.Query,
	Mutation: mutation.Mutation,
})

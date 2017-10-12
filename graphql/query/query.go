package query

import (
	"github.com/graphql-go/graphql"

	. "go-todo-api/models"
	. "go-todo-api/graphql/types"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"product": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
        product := ShowProduct(params.Args["id"].(string))
				return product, nil
			},
		},

		"products": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Description: "List products",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return AllProduct(), nil
			},
		},
	},
})

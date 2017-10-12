package mutation

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"

	. "go-todo-api/models"
	. "go-todo-api/graphql/types"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		"createProduct": &graphql.Field{
			Type: ProductType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
        res := CreateProduct(Product{Name: name})

				return res, nil
			},
		},

		"updateProduct": &graphql.Field{
			Type: ProductType,
			Description: "Update existing product",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(string)
				name, _ := params.Args["name"].(string)

        UpdateProduct(id, bson.M{name: name})

        return Product{Id: id, Name: name}, nil
			},
		},

	},
})

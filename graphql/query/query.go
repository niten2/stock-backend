package query

import (
  // "fmt"
  // "os"
  // "reflect"
	// "time"
	// "math/rand"
  // "os"
	"github.com/graphql-go/graphql"
  // "github.com/fatih/structs"
  // "gopkg.in/mgo.v2"

	. "go-todo-api/models"
	// "go-todo-api/db"
)

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{ Type: graphql.String },
		"name": &graphql.Field{ Type: graphql.String },
	},
})

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// "product": &graphql.Field{
		// 	Type: productType,

		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        // // todo := Product{ID: "xcccccca", Name: "sdfsdfsdfdsf"}

        // // product := Product{doc.Id, doc.Name}

		// 		// return product, nil
		// 		return nil, nil
		// 	},
		// },

		"products": &graphql.Field{
			Type: graphql.NewList(productType),
			Description: "List products",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return AllProduct(), nil
			},
		},
	},
})








// // curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
// "todo": &graphql.Field{
  // Type:        todoType,
  // Description: "Get single todo",
  // Args: graphql.FieldConfigArgument{
    // "id": &graphql.ArgumentConfig{
      // Type: graphql.String,
    // },
  // },
  // Resolve: func(params graphql.ResolveParams) (interface{}, error) {

    // idQuery, isOK := params.Args["id"].(string)
    // if isOK {
      // // Search for el with id
      // for _, todo := range TodoList {
        // if todo.ID == idQuery {
          // return todo, nil
        // }
      // }
    // }

    // return Todo{}, nil
  // },
// },

// "lastTodo": &graphql.Field{
  // Type:        todoType,
  // Description: "Last todo added",
  // Resolve: func(params graphql.ResolveParams) (interface{}, error) {
    // return TodoList[len(TodoList)-1], nil
  // },
// },


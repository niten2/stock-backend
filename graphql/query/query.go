package query

import (
  // "fmt"
  // "reflect"
	// "time"
	// "math/rand"
	"github.com/graphql-go/graphql"
  // "gopkg.in/mgo.v2"
  // "os"
)

// var TodoList []Todo


// type Todo struct {
// 	ID   string `json:"id"`
// 	Name string `json:"text"`
// 	Title string `json:"text"`
// }


// var todoType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Todo",

// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{Type: graphql.String},
// 		"name": &graphql.Field{Type: graphql.String},
// 		"title": &graphql.Field{Type: graphql.String},
// 	},
// })


// type Product struct {
// 	ID   string `json:"id"`
// 	Name string `json:"text"`
// }

// var productType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "product",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"name": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 		},
// 	},
// )

// type ProductDocument struct {
// 	Id string `bson:"_id,omitempty"`
// 	Name string
//   Phone string
// }

// type Product struct {
//   Id string
//   Name string
//   Phone string
// }


// func init() {
// 	todo1 := Todo{ID: "a", Title: "A 3333todo not to forget"}
// 	// todo2 := Todo{ID: "b", Name2: "This is the most important"}
// 	// todo3 := Todo{ID: "c", Name2: "Please do this or else"}
//   // fmt.Println(todo1)

// 	// TodoList = append(TodoList, todo1, todo2, todo3)
// 	TodoList = append(TodoList, todo1)
//   // fmt.Println(TodoList)


// 	rand.Seed(time.Now().UnixNano())
// }


var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		// "todo": &graphql.Field{
		// 	Type: todoType,

		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        // todo := Todo{ID: "xcccccca", NameTitle: "sdfsdfsdfdsf"}
        // // fmt.Println(todo.Title)

        // return todo, nil
		// 	},
		// },

		// "product": &graphql.Field{
		// 	Type: productType,

		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        // // todo := Product{ID: "xcccccca", Name: "sdfsdfsdfdsf"}

        // // product := Product{doc.Id, doc.Name}

		// 		// return product, nil
		// 		return nil, nil
		// 	},
		// },


		// "products": &graphql.Field{
		// 	Type:        graphql.NewList(productType),
		// 	Description: "List of todos",

		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        // // fmt.Println(TodoList)

        // todo1 := Product{ID: "xcccccca", Name: "sdfsdfsdfdsf"}
        // todo2 := Product{ID: "xcccccca", Name: "sdfsdfsdfdsf"}

		// 		// return TodoList, nil
		// 		return todo, nil
		// 	},
		// },
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


package mutation

import (
  "fmt"
	"github.com/graphql-go/graphql"
	// "github.com/graphql-go/relay"
	// "golang.org/x/net/context"
	// "gopkg.in/mgo.v2/bson"

	. "go-todo-api/models"
	. "go-todo-api/schema/types"
	. "go-todo-api/schema/inputs"
)

// var createProduct = relay.MutationWithClientMutationID(relay.MutationConfig{
//   Name: "createProduct",
//   InputFields: graphql.InputObjectConfigFieldMap{
//     "name": &graphql.InputObjectFieldConfig{
//       Type: graphql.NewNonNull(graphql.String),
//     },
//   },
//   // OutputFields: graphql.Fields{
//   //   "id": &graphql.Field{
//   //     Type: ProductType,
//   //     Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//   //       fmt.Println(p)

//   //       // if payload, ok := p.Source.(map[string]interface{}); ok {
//   //       // 	return GetShip(payload["shipId"].(string)), nil
//   //       // }
//   //       return nil, nil
//   //     },
//   //   },
//   // },
//   MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {

//     fmt.Println(111111111)

//     name := inputMap["name"].(string)
//     res := CreateProduct(Product{Name: name})

//     fmt.Println(res.Id)
//     fmt.Println(res.Name)

//     return map[string]interface{}{
//       "id": res.Id,
//       "name": res.Name,
//     }, nil

//     // // `inputMap` is a map with keys/fields as specified in `InputFields`
//     // // Note, that these fields were specified as non-nullables, so we can assume that it exists.
//     // shipName := inputMap["shipName"].(string)
//     // factionId := inputMap["factionId"].(string)

//     // // This mutation involves us creating (introducing) a new ship
//     // newShip := CreateShip(shipName, factionId)
//     // // return payload
//     // return map[string]interface{}{
//     // 	"shipId":    newShip.ID,
//     // 	"factionId": factionId,
//     // }, nil
//   },
// })

func init () {

  fmt.Println(ProductInput)
}


var createProduct = &graphql.Field{
  Type: ProductType,
  Description: "Create new product",
  Args: graphql.FieldConfigArgument{
  	"input": &graphql.ArgumentConfig{
  		Type: ProductInput,
  	},
  },
  Resolve: func(params graphql.ResolveParams) (interface{}, error) {
    name, _ := params.Args["name"].(string)
    res := CreateProduct(Product{Name: name})
    return res, nil
  },
}



var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

    "createProduct": createProduct,



		// "createProduct": &graphql.Field{
		// 	Type: ProductType,
		// 	Description: "Create new product",
		// 	// Args: graphql.FieldConfigArgument{
		// 	// 	// "name": &graphql.ArgumentConfig{
		// 	// 	// 	Type: graphql.NewNonNull(graphql.String),
		// 	// 	// },
		// 	// 	"input": &graphql.ArgumentConfig{
		// 	// 		Type: graphql.NewNonNull(graphql.String),
		// 	// 	},
		// 	// },
		// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// 		name, _ := params.Args["name"].(string)
        // res := CreateProduct(Product{Name: name})
		// 		return res, nil
		// 	},
		// },

		// "updateProduct": &graphql.Field{
		// 	Type: ProductType,
		// 	Description: "Update existing product",
		// 	Args: graphql.FieldConfigArgument{
		// 		"id": &graphql.ArgumentConfig{
		// 			Type: graphql.NewNonNull(graphql.String),
		// 		},
		// 		"name": &graphql.ArgumentConfig{
		// 			Type: graphql.NewNonNull(graphql.String),
		// 		},
		// 	},
		// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// 		id, _ := params.Args["id"].(string)
		// 		name, _ := params.Args["name"].(string)

        // UpdateProduct(id, bson.M{name: name})

        // return Product{Id: id, Name: name}, nil
		// 	},
		// },

	},
})

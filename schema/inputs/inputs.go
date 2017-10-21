package inputs

import (
  // "fmt"
  // "os"
  // "reflect"
	// "time"
	// "math/rand"
  // "os"
  // "github.com/fatih/structs"
  // "gopkg.in/mgo.v2"
	"github.com/graphql-go/graphql"

	// . "go-todo-api/models"
	// "go-todo-api/db"
)

// var ProductInput = graphql.NewInputObject(graphql.InputObjectConfig{
// 	Name: "ProductInput",
// 	Fields: graphql.Fields{
// 		"name": &graphql.Field{ Type: graphql.String },
// 	},
// })

var ProductInput = graphql.NewInputObject(
    graphql.InputObjectConfig{
        Name: "ProductInput",
        Fields: graphql.InputObjectConfigFieldMap{
            "name": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
        },
    },
)

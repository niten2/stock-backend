package types

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

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{ Type: graphql.String },
		"name": &graphql.Field{ Type: graphql.String },
	},
})


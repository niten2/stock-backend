package testing

import (
  // "fmt"
  "testing"
	"reflect"
  "github.com/graphql-go/graphql"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"

  "go-todo-api/db"
  . "go-todo-api/models"
  . "go-todo-api/graphql"
)

func init() {
  _ = godotenv.Load("../../.env.test")
	db.Connect()
}

func TestQueryProducts(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("AllProduct", t, func() {
    CreateProduct(Product{Id: "1", Name: "test1"})
    CreateProduct(Product{Id: "2", Name: "test2"})

    query := "query { products { id name } }"
    res := ExecuteQuery(query)

    expected := &graphql.Result{
      Data: map[string]interface{}{
        "products": []interface{}{
          map[string]interface{}{
            "id":   "1",
            "name": "test1",
          },
          map[string]interface{}{
            "id":   "2",
            "name": "test2",
          },
        },
      },
    }

    eq := reflect.DeepEqual(res, expected)
    So(eq, ShouldBeTrue)
  })
}

func TestQueryProduct(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("product", t, func() {
    CreateProduct(Product{Id: "1", Name: "test1"})

    query := `query {
      product(id: "1") {
        id
        name
      }
    }`

    res := ExecuteQuery(query)

    expected := &graphql.Result{
      Data: map[string]interface{}{
        "product": map[string]interface{}{
          "id":   "1",
          "name": "test1",
        },
      },
    }

    So(reflect.DeepEqual(res, expected), ShouldBeTrue)
  })
}


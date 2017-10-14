package testing

import (
	// "reflect"
  // "github.com/mailgun/godebug"
  "testing"
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

func TestCreateProducts(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("createProduct", t, func() {
    query := `mutation {
      createProduct(name: "name") {
        id
        name
      }
    }`
    res := ExecuteQuery(query)

    id := res.Data.(map[string]interface{})["createProduct"].(map[string]interface{})["id"].(string)

    product := ShowProduct(id)
    So(product.Name, ShouldEqual, "name")
  })
}

func TestUpdateProduct(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("UpdateProduct", t, func() {
    CreateProduct(Product{Id: "1", Name: "test1"})

    query := `mutation {
      updateProduct(id: "1", name: "name") {
        id
        name
      }
    }`

    res := ExecuteQuery(query)
    id := res.Data.(map[string]interface{})["updateProduct"].(map[string]interface{})["id"].(string)
    product := ShowProduct(id)

    So(product.Name, ShouldEqual, "name")
  })
}


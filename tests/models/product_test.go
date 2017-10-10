package testing

import (
  // "fmt"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  _ "github.com/joho/godotenv/autoload"
  "github.com/joho/godotenv"

  . "go-todo-api/models"
  . "go-todo-api/db"
	"gopkg.in/mgo.v2/bson"
)

func TestSomething(t *testing.T) {
  _ = godotenv.Load("../../.env.test")
  db := ConnectDb().C("Products")
  db.RemoveAll(nil)

  Convey("AllProduct", t, func() {

    db.RemoveAll(nil)

    product1 := Product{Name: "test1"}
    product2 := Product{Name: "test2"}
    db.Insert(product1)
    db.Insert(product2)

    res := AllProduct()

    So(res, ShouldHaveLength, 2)
  })

  Convey("CreateProduct", t, func() {
    res := CreateProduct(Product{Name: "test"})

    So(res, ShouldEqual, true)
  })

  Convey("ShowProduct", t, func() {
    db.Insert(bson.M{"_id": "1234", "name": "test"})
    res := ShowProduct("1234")

    So(res.Name, ShouldEqual, "test")
  })
}

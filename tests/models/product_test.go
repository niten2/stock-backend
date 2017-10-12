package testing

import (
  "testing"
  "github.com/joho/godotenv"
  . "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"

  "go-todo-api/db"
  . "go-todo-api/models"
)

func init() {
  _ = godotenv.Load("../../.env.test")
	db.Connect()
}

func TestSomething(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("AllProduct", t, func() {
    CreateProduct(Product{Name: "test1"})
    CreateProduct(Product{Name: "test2"})

    So(AllProduct(), ShouldHaveLength, 2)
  })

  Convey("CreateProduct", t, func() {
    res := CreateProduct(Product{Name: "test"})

    So(res.Name, ShouldEqual, "test")
    So(res.Id, ShouldNotBeBlank)
  })

  Convey("ShowProduct", t, func() {
    DbProduct.Insert(bson.M{"_id": "1234", "name": "test"})
    res := ShowProduct("1234")

    So(res.Name, ShouldEqual, "test")
  })

  Convey("DeleteProduct", t, func() {
    DbProduct.Insert(bson.M{"_id": "1234", "name": "test"})
    res := DeleteProduct("1234")

    response := ShowProduct("1234")

    So(response.Name, ShouldEqual, "")
    So(res, ShouldEqual, true)
  })

  Convey("UpdateProduct", t, func() {
    DbProduct.Insert(bson.M{"_id": "1234", "name": "test"})
    UpdateProduct("1234", bson.M{"name": "new test"})

    response := ShowProduct("1234")

    So(response.Name, ShouldEqual, "new test")
  })

}

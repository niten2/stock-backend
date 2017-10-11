package product

import (
  "fmt"
	// "log"
  // "os"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "go-todo-api/db"
)

// type ProductDocument struct {
// 	Id string `bson:"_id,omitempty"`
// 	Name string
// }

type Product struct {
  ID bson.ObjectId `bson:"_id,omitempty"`
	Name string `json: "username"`
	// Age	int `json: "age"`
	// LastUpdated time.Time
}

func AllProduct() []Product {
  db := ConnectDb().C("Products")

	var products []Product

  err := db.Find(bson.M{}).All(&products)
  if err != nil {
    fmt.Println(err)
  }

	return products
}

func CreateProduct(product Product) bool {
  db := ConnectDb().C("Products")
  err := db.Insert(product)

  if err != nil {
    fmt.Println(err)
    return false
  }

	return true
}


func ShowProduct(id string) Product {
  db := ConnectDb().C("Products")
	product := Product{}

  err := db.FindId(id).One(&product)
  if err != nil {
    fmt.Println(err)
  }

	return product
}

func DeleteProduct(id string) bool {
  db := ConnectDb().C("Products")

  err := db.RemoveId(id)
  if err != nil {
    fmt.Println(err)
    return false
  }

	return true
}

func UpdateProduct(id string, attributes bson.M) bool {
  db := ConnectDb().C("Products")

  err := db.Update(bson.M{"_id": id}, bson.M{"$set": attributes})
  if err != nil {
    fmt.Println(err)
    return false
  }

	return true
}

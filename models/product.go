package product

import (
  "fmt"
	// "log"
  // "os"
	// "gopkg.in/mgo.v2"
  // "fmt"
	"gopkg.in/mgo.v2/bson"

	"go-todo-api/db"
)

type Product struct {
	// Name string `json: "name"`
	// ID string `json:"id"`
  // ID bson.ObjectId `bson:"_id,omitempty"`
  // Id string `json:"_id" bson:"_id,omitempty"`
  // Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`

  // Id string `json:"id" bson:"_id,omitempty"`
  Id string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	// Info string `json: "info"`
	// LastUpdated time.Time
}

func AllProduct() []Product {
  DbProduct := db.Db.C("products")

	var products []Product

  err := DbProduct.Find(bson.M{}).All(&products)
  if err != nil {
    fmt.Println(err)
  }

	return products
}

func CreateProduct(product Product) Product {
  DbProduct := db.Db.C("products")

  if product.Id == "" {
    product.Id = bson.NewObjectId().Hex()
  }

  err := DbProduct.Insert(product)
  if err != nil {
    fmt.Println(err)
  }

	return product
}

func ShowProduct(id string) Product {
  DbProduct := db.Db.C("products")

	product := Product{}

  err := DbProduct.FindId(id).One(&product)
  if err != nil {
    fmt.Println(err)
  }

	return product
}

func DeleteProduct(id string) bool {
  DbProduct := db.Db.C("products")

  err := DbProduct.RemoveId(id)
  if err != nil {
    fmt.Println(err)
    return false
  }

	return true
}

func UpdateProduct(id string, attributes bson.M) bool {
  DbProduct := db.Db.C("products")

  err := DbProduct.Update(bson.M{"_id": id}, bson.M{"$set": attributes})
  if err != nil {
    fmt.Println(err)
    return false
  }

	return true
}

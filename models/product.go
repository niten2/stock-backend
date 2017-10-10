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
}


// Profile - is the memory representation of one user profile
// type Profile struct {
// 	Name 			string			`json: "username"`
// 	Password 		string			`json: "password"`
// 	Age 			int				`json: "age"`
// 	LastUpdated     time.Time
// }

// func AllProduct() []Product {
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

// // DeleteProfile - Deletes the profile in the Profiles Collection with name equal to the id parameter (id == name)
// func DeleteProfile(id string) bool {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		log.Println("Could not connect to mongo: ", err.Error())
// 		return false
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("ProfileService").C("Profiles")
// 	err = c.RemoveId(id)

// 	return true
// }

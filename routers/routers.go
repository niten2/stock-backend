package routers

import (
	"net/http"
  "github.com/go-martini/martini"
  "gopkg.in/mgo.v2"
  "github.com/codegangsta/martini-contrib/render"
)

type ProductDocument struct {
	Id string `bson:"_id,omitempty"`
	Name string
  Phone string
}

type Product struct {
  Id string
  Name string
  Phone string
}

func ProductsHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {

	productDocuments := []Product{}
	productCollection := db.C("products")
	productCollection.Find(nil).All(&productDocuments)

	products := []Product{}
	for _, doc := range productDocuments {
		product := Product{doc.Id, doc.Name, doc.Phone}
		products = append(products, product)
	}

  rnd.JSON(200, products)
}

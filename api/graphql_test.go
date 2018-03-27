package main_test

import (
	"fmt"
	// "encoding/json"
	// "net/http"
	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"

	"go-todo-api/db"
	. "go-todo-api/models"
	"go-todo-api/routers"
)

func init() {
	_ = godotenv.Load("../../.env.test")
	db.Connect()
}

func TestQueryGraphql(t *testing.T) {
	Convey("POST /v1", t, func() {

		DbProduct := db.Db.C("products")
		DbProduct.RemoveAll(nil)

		Convey("products", t, func() {
			CreateProduct(Product{Id: "1", Name: "test1"})
			CreateProduct(Product{Id: "2", Name: "test2"})

			params := `query {
        products {
          id
          name
        }
      }`

			// req := httptest.NewRequest("GET", "/v1", nil)
			req := httptest.NewRequest("GET", "/v1", params)
			w := httptest.NewRecorder()

			routers.Handlers().ServeHTTP(w, req)
			fmt.Println(w)

			// expected := &graphql.Result{
			//   Data: map[string]interface{}{
			//     "products": []interface{}{
			//       map[string]interface{}{
			//         "id":   "1",
			//         "name": "test1",
			//       },
			//       map[string]interface{}{
			//         "id":   "2",
			//         "name": "test2",
			//       },
			//     },
			//   },
			// }

			// eq := reflect.DeepEqual(res, expected)
			// So(eq, ShouldBeTrue)
		})

	})

	// Convey("GET /v1", t, func() {

	//   DbProduct := db.Db.C("products")
	//   DbProduct.RemoveAll(nil)

	//   // Convey("products", t, func() {
	//     CreateProduct(Product{Id: "1", Name: "test1"})
	//     CreateProduct(Product{Id: "2", Name: "test2"})

	//     // query := `/v1?query {
	//     //   products {
	//     //     id
	//     //     name
	//     //   }
	//     // }`

	//     query := `/v1?query{ products { id name } }`

	//     // fmt.Println("/v1" + query)

	//     // req := httptest.NewRequest("GET", "/v1", nil)
	//     req := httptest.NewRequest("GET", query, nil)
	//     w := httptest.NewRecorder()

	//     routers.Handlers().ServeHTTP(w, req)
	//     fmt.Println(w)

	//     // expected := &graphql.Result{
	//     //   Data: map[string]interface{}{
	//     //     "products": []interface{}{
	//     //       map[string]interface{}{
	//     //         "id":   "1",
	//     //         "name": "test1",
	//     //       },
	//     //       map[string]interface{}{
	//     //         "id":   "2",
	//     //         "name": "test2",
	//     //       },
	//     //     },
	//     //   },
	//     // }

	//     // eq := reflect.DeepEqual(res, expected)
	//     // So(eq, ShouldBeTrue)
	//   // })

	// })
}

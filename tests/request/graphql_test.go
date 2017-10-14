package main_test

import (
  "fmt"
  // "encoding/json"
  "net/http/httptest"
  "net/http"
  // "bytes"
  // "io"
  // "net/url"
  "github.com/joho/godotenv"
  "go-todo-api/db"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  . "go-todo-api/models"
  "go-todo-api/routers"
)

func init() {
  _ = godotenv.Load("../../.env.test")
	db.Connect()
}


// func ExecuteQuery(query string) httptest.NewRequest {
//   req := httptest.NewRequest("GET", "/v1", query)
//   w := httptest.NewRecorder()
//   routers.Handlers().ServeHTTP(w, req)
//   return w
// }

func TestQueryGraphql(t *testing.T) {
  Convey("GET /v1", t, func() {

    DbProduct := db.Db.C("products")
    DbProduct.RemoveAll(nil)

    CreateProduct(Product{Id: "1", Name: "test1"})
    CreateProduct(Product{Id: "2", Name: "test2"})

    query := `{
      products {
        id
        name
      }
    }`


    req, _ := http.NewRequest("POST", "/v1", nil) // <-- URL-encoded payload
    // req, _ := http.NewRequest("POST", "/v1?query={products{id,name}}", nil) // <-- URL-encoded payload

    q := req.URL.Query()
    q.Add("query", query)
    req.URL.RawQuery = q.Encode()

    // fmt.Println(req.URL.String())


    // r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    // r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    // r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    // resp, _ := client.Do(r)
    // fmt.Println(resp)

    // fmt.Println(r)


      // req := httptest.NewRequest("GET", "/v1", buf)
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



  // user := &User{Name: "Test user"}
  // data, err := json.Marshal(``)
  // So(err, ShouldBeNil)
  // buf := bytes.NewBuffer(data)



   // fmt.Println(w.Result())
   // So(w.Body.String(), ShouldContain, "Must provide an operation.")
  })
}

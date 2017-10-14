goconvey
append(TodoList, todo1, todo2, todo3)
rand.Seed(time.Now().UnixNano())



  // s3Bucket := os.Getenv("S3_BUCKET")
  // secretKey := os.Getenv("SECRET_KEY")

	// fmt.Println(s3Bucket)
	// fmt.Println(secretKey)


curl -g 'http://localhost:8000/graphql?query={todo(id:\"b\"){id,text,done}}'

curl -g 'http://localhost:8000/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'


gin -p 4000 --appPort 4001 run main.go



fresh

http://localhost:8000/graphql?query={todoList{id,text,done}}

http://localhost:8000/graphql?query={todoList{id,name}}

http://localhost:8000/graphql?query={products{id,name}}

curl -g 'http://localhost:8000/graphql?query={products{id,name}}'

curl -g 'http://localhost:8000/graphql?query={todoList{id,title}}'


curl -g 'http://localhost:8000/graphql?query={todo{id,title, name}}'


curl -g 'http://localhost:8000/graphql?query={todoList{id,text,done}}'

curl -g 'http://localhost:8000/graphql?query=mutation+_{createTodo(text:"11111My+new+todo"){id,text,done}}'


// curl -g 'http://localhost:8000/graphql?query={todo(id:\"b\"){id,text,done}}'
// fmt.Println("Get single todo: curl -g 'http://localhost:8080/graphql?query={todo(id:\"b\"){id,text,done}}'")
// fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
// fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")
// fmt.Println("Load todo list: curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'")
// fmt.Println("Access the web app via browser at 'http://localhost:8080'")
//
//

// // main.go (web-server)
// // +build integration

package graphql_test

func main() {
}
// import (
//   "fmt"
//   // "io"
//   // "io/ioutil"
//   // "log"
//   // "net/http"
//   "net/http/httptest"
//   "testing"
// 	// "reflect"
//   // "github.com/graphql-go/graphql"
//   "github.com/joho/godotenv"
//   // . "github.com/smartystreets/goconvey/convey"
//   // "encoding/json"
//   // "github.com/mailgun/godebug"

//   // "go-todo-api/db"
//   // . "go-todo-api/models"
// )

// func init() {
//   _ = godotenv.Load("../../.env.test")
// 	// db.Connect()
// }

// // func TestWrongUsernamePassword(t *testing.T) {
// // 	if a.Login("user1", "wrongpassword").Token != "" {
// // 		t.Fail()
// // 	}
// // }

// // var (
// //   server   *httptest.Server
// //   reader   io.Reader
// //   usersUrl string
// // )

// // func init() {
// //   server = httptest.NewServer(Handlers())
// //   usersUrl = fmt.Sprintf("%s/users", server.URL)
// // }

// func TestQueryProductstest(t *testing.T) {
//  // req, _ := http.NewRequest("GET", "http://localhost/", nil)
//  // fmt.Println(req)

//    // res := http.Get("127.0.0.1:8001")
//    // fmt.Println(res)

//    // res := http.NewRequest("GET", "/", nil)

//    req := httptest.NewRequest("GET", "/", nil)
//    w := httptest.NewRecorder()
//    // handler(w, req)
//    fmt.Println(Handlers())


//    fmt.Println(w)


// 	 // req, _ := http.NewRequest("GET", "/", nil)

//    // fmt.Println(req)


//    // userJson := `{"username": "dennis", "balance": 200}`

//   // reader = strings.NewReader(userJson) //Convert string to reader

//   // request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

//   // res, err := http.DefaultClient.Do(request)

//   // if err != nil {
//    //    t.Error(err) //Something is wrong while sending request
//   // }

//   // if res.StatusCode != 201 {
//    //    t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
//   // }




//   // ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//   //   w.WriteHeader(http.StatusServiceUnavailable)
//   // }))
//   // defer ts.Close()

//   // fmt.Println(ts)


//   // nsqdUrl := ts.URL
//   // // fmt.Println(nsqdUrl)

//   // fmt.Println(1111111111111111)


//   // err := Publish(nsqdUrl, "hello")
//   // if err == nil {
//   //   t.Errorf("Publish() didn’t return an error")
//   // }


//   // ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//   //   fmt.Fprintln(w, "Hello, client")
//   // }))
//   // defer ts.Close()

//   // client := ts.Client()
//   // res, err := client.Get(ts.URL)
//   // if err != nil {
//   //   log.Fatal(err)
//   // }

//   // greeting, err := ioutil.ReadAll(res.Body)
//   // res.Body.Close()
//   // if err != nil {
//   //   log.Fatal(err)
//   // }

//   // fmt.Printf("%s", greeting)
//   // // Output: Hello, client
// }



// // func TestCreateProducts(t *testing.T) {
// //   DbProduct := db.Db.C("products")
// //   DbProduct.RemoveAll(nil)

// //   Convey("createProduct", t, func() {
// //     query := `mutation {
// //       createProduct(name: "name") {
// //         id
// //         name
// //       }
// //     }`
// //     res := ExecuteQuery(query)

// //     id := res.Data.(map[string]interface{})["createProduct"].(map[string]interface{})["id"].(string)

// //     product := ShowProduct(id)
// //     So(product.Name, ShouldEqual, "name")
// //   })
// // }

// // func TestUpdateProduct(t *testing.T) {
// //   DbProduct := db.Db.C("products")
// //   DbProduct.RemoveAll(nil)

// //   Convey("UpdateProduct", t, func() {
// //     CreateProduct(Product{Id: "1", Name: "test1"})

// //     query := `mutation {
// //       updateProduct(id: "1", name: "name") {
// //         id
// //         name
// //       }
// //     }`

// //     res := ExecuteQuery(query)
// //     id := res.Data.(map[string]interface{})["updateProduct"].(map[string]interface{})["id"].(string)
// //     product := ShowProduct(id)

// //     So(product.Name, ShouldEqual, "name")
// //   })
// // }


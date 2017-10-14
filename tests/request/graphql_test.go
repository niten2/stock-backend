package main_test

import (
  // "fmt"
  // "encoding/json"
  "net/http/httptest"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "go-todo-api/routers"
)

func TestQueryGraphql(t *testing.T) {
  Convey("GET /v1", t, func() {
   req := httptest.NewRequest("GET", "/v1", nil)
   w := httptest.NewRecorder()
   routers.Handlers().ServeHTTP(w, req)

   // fmt.Println(w.Result())
   // So(w.Body.String(), ShouldContain, "Must provide an operation.")
  })
}

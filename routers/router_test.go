package routers_test

import (
  "fmt"
  "testing"
  // "net/http/httptest"
  . "github.com/smartystreets/goconvey/convey"

)

func TestQueryMain(t *testing.T) {
  Convey("GET /", t, func() {

    fmt.Println(11111)

   // req := httptest.NewRequest("GET", "/", nil)
   // w := httptest.NewRecorder()

   // routers.Handlers().ServeHTTP(w, req)

   // expect := `{
   //   "graphql": "/v1",
   //   "service": "go-todo",
   //   "version": "v1"
   // }`

   // So(w.Body.String(), ShouldEqual, expect)
  })
}

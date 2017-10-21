package main_test

import (
  "net/http/httptest"
  "testing"
  . "github.com/smartystreets/goconvey/convey"

  "go-todo-api/routers"
)

func TestQueryMain(t *testing.T) {
  Convey("GET /", t, func() {

   req := httptest.NewRequest("GET", "/", nil)
   w := httptest.NewRecorder()

   routers.Handlers().ServeHTTP(w, req)

   expect := `{
     "graphql": "/v1",
     "service": "go-todo",
     "version": "v1"
   }`

   So(w.Body.String(), ShouldEqual, expect)
  })
}

// package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"runtime"
// 	"path/filepath"
// 	_ "todo-api/routers"

// 	"github.com/astaxie/beego"
// 	. "github.com/smartystreets/goconvey/convey"
// )

// func init() {
// 	_, file, _, _ := runtime.Caller(1)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
// 	beego.TestBeegoInit(apppath)
// }

// // TestGet is a sample to run an endpoint test
// func TestGet(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/object", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 	        })
// 	        Convey("The Result Should Not Be Empty", func() {
// 	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 	        })
// 	})
// }

package testing

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
    Convey("Given some integer with a starting value", t, func() {
        x := 1

        Convey("When the integer is incremented", func() {
            x++

            Convey("The value should be greater by one", func() {
                So(x, ShouldEqual, 4)
            })
        })
    })
}

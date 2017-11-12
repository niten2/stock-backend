package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Home")
	// addGraphQLHandler()
}

// func addGraphQLHandler() {
// 	beego.Router("/query", &controllers.GraphqlController{}, "post:Query")
// 	beego.Router("/graphql", &controllers.GraphqlController{}, "get:Index")
// }

package api

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Home() {
	this.Ctx.WriteString("hello world")
}

package controllers

import (
	// "todo-api/models"
	// "encoding/json"

  "github.com/astaxie/beego"
)

type TestController struct {
  beego.Controller
}

func (this *TestController) GetAll() {
  // this.Data["Website"] = "beego.me"
  // this.Data["Email"] = "astaxie@gmail.com"
  // this.TplNames = "index.tpl"

  this.Data["json"] = "dsfsdfdsfsdf"
	this.ServeJSON()
}

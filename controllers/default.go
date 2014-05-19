package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
}

func (this *MainController) Get() {
	this.Ctx.Output.Body([]byte("Docker Registry"))
}

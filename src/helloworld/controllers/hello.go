package controllers

import "github.com/astaxie/beego"

type HelloControllers struct {
	beego.Controller
}

func (hello * HelloControllers) Get() {
	hello.Ctx.WriteString("hello world")
}

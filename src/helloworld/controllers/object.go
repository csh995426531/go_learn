package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

func (this *ObjectController) GetFunc() {
	params := this.Ctx.Input.Params()

	for k,v := range params{
		fmt.Println(k, "=", v)
	}

	this.Ctx.WriteString(params[":splat"])
}

package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {
	
}

// @router /all/:key [get]
func (this *CMSController) AllBlock()  {

}

package routers

import (
	"helloworld/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloControllers{})
    beego.AutoRouter(&controllers.ObjectController{})
	beego.Include(&controllers.CMSController{})
}

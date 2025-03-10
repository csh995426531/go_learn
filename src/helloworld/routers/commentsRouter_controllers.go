package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["helloworld/controllers:CMSController"] = append(beego.GlobalControllerRouter["helloworld/controllers:CMSController"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: `/all/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["helloworld/controllers:CMSController"] = append(beego.GlobalControllerRouter["helloworld/controllers:CMSController"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: `/staticblock/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

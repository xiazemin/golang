package routers

import (
	"helloBee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/helloBee", &controllers.MainController{}, "get:HelloSitepoint")
}

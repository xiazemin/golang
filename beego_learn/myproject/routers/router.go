package routers

import (
	"github.com/xiazemin/beego_learn/myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

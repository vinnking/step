package routers

import (
	"step/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 首页
    beego.Router("/", &controllers.MainController{})
	
	// 用户模块
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
	beego.Router("/user/create", &controllers.UserController{}, "GET,POST:Create")
}

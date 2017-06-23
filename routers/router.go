package routers

import (
	"github.com/astaxie/beego"
	"step/controllers"
)

func init() {
	// 首页
	beego.Router("/", &controllers.MainController{})

	// 用户模块
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
	beego.Router("/user/create", &controllers.UserController{}, "GET,POST:Create")
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{}, "GET,POST:View")
}

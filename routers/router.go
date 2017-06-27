package routers

import (
	"github.com/astaxie/beego"

	"step/controllers"
)

func init() {
	// 首页
	beego.Router("/", &controllers.MainController{})

	// 用户模块, 列表, 添加, 明细, 更新, 停用, 删除
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
	beego.Router("/user/create", &controllers.UserController{}, "GET,POST:Create")
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{}, "GET,POST:View")
	beego.Router("/user/update/:id([0-9]+)", &controllers.UserController{}, "POST:Update")
	beego.Router("/user/stop/:id([0-9]+)", &controllers.UserController{}, "POST:Stop")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "POST:Delete")
}

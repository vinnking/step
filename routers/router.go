package routers

import (
	"github.com/astaxie/beego"

	"step/controllers"
)

func init() {
	// 首页
	beego.Router("/?:id([0-9]+)", &controllers.MainController{})

	// 用户模块, 列表, 添加, 明细, 更新, 停用
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
	beego.Router("/user/create", &controllers.UserController{}, "GET,POST:Create")
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{}, "GET,POST:View")
	beego.Router("/user/update/:id([0-9]+)", &controllers.UserController{}, "GET,POST:Update")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")
	
	// 引用模块
	beego.Router("/quote", &controllers.QuoteController{}, "GET:Index")
	beego.Router("/quote/create", &controllers.QuoteController{}, "GET,POST:Create")
	beego.Router("/quote/:id([0-9]+)", &controllers.QuoteController{}, "GET,POST:View")
	beego.Router("/quote/update/:id([0-9]+)", &controllers.QuoteController{}, "GET,POST:Update")
	beego.Router("/quote/delete/:id([0-9]+)", &controllers.QuoteController{}, "GET:Delete")
}

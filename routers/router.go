package routers

import (
	"github.com/astaxie/beego"

	"step/controllers"
)

func init() {
	// 首页
	beego.Router("/?:id([0-9]+)", &controllers.MainController{})

	// 用户, 列表, 添加, 明细, 更新, 停用
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
	beego.Router("/user/create", &controllers.UserController{}, "GET,POST:Create")
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{}, "GET,POST:View")
	beego.Router("/user/update/:id([0-9]+)", &controllers.UserController{}, "GET,POST:Update")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")

	// 引用
	beego.Router("/quote", &controllers.QuoteController{}, "GET:Index")
	beego.Router("/quote/create", &controllers.QuoteController{}, "GET,POST:Create")
	beego.Router("/quote/:id([0-9]+)", &controllers.QuoteController{}, "GET,POST:View")
	beego.Router("/quote/update/:id([0-9]+)", &controllers.QuoteController{}, "GET,POST:Update")
	beego.Router("/quote/delete/:id([0-9]+)", &controllers.QuoteController{}, "GET:Delete")

	// 标签
	beego.Router("/label", &controllers.LabelController{}, "GET:Index")
	beego.Router("/label/create", &controllers.LabelController{}, "GET,POST:Create")
	beego.Router("/label/:id([0-9]+)", &controllers.LabelController{}, "GET,POST:View")
	beego.Router("/label/update/:id([0-9]+)", &controllers.LabelController{}, "GET,POST:Update")
	beego.Router("/label/delete/:id([0-9]+)", &controllers.LabelController{}, "GET:Delete")

	// 友情链接
	beego.Router("/link", &controllers.LinkController{}, "GET:Index")
	beego.Router("/link/create", &controllers.LinkController{}, "GET,POST:Create")
	beego.Router("/link/:id([0-9]+)", &controllers.LinkController{}, "GET,POST:View")
	beego.Router("/link/update/:id([0-9]+)", &controllers.LinkController{}, "GET,POST:Update")
	beego.Router("/link/delete/:id([0-9]+)", &controllers.LinkController{}, "GET:Delete")
}

package routers

import (
	"github.com/astaxie/beego"

	"step/controllers"
)

func init() {
	// 首页, 详情页, 标签
	beego.Router("/?:id([0-9]+)", &controllers.MainController{})
	beego.Router("/view/:id([0-9]+)", &controllers.MainController{}, "GET:View")
	beego.Router("/labels/:id([0-9]+)", &controllers.MainController{}, "GET:Label")
	
	// 管理员登陆和退出
	beego.Router("/auth/login", &controllers.AuthController{}, "GET,POST:Login")
	beego.Router("/auth/logout", &controllers.AuthController{}, "GET,POST:Logout")

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
	
	// 文章
	beego.Router("/post", &controllers.PostController{}, "GET:Index")
	beego.Router("/post/create", &controllers.PostController{}, "GET,POST:Create")
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "GET,POST:View")
	beego.Router("/post/update/:id([0-9]+)", &controllers.PostController{}, "GET,POST:Update")
	beego.Router("/post/delete/:id([0-9]+)", &controllers.PostController{}, "GET:Delete")
}

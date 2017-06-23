// 用户
package controllers

import (
	"github.com/astaxie/beego"
	"step/models"
	"fmt"
	"strings"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

// 用户列表
func (u *UserController) Index() {
	u.Layout = "base.html"
	u.TplName = "user/index.html"
}

// 添加用户
func (u *UserController) Create() {
	beego.ReadFromRequest(&u.Controller)
	if u.Ctx.Request.Method == "POST" {
		nickname := u.Input().Get("username")
		email := u.Input().Get("email")
		password := u.Input().Get("password")
		role := u.Input().Get("role")
		// 用户名成不能为空
		if strings.TrimSpace(nickname) == "" {
			u.Redirect("/user/create", 302)
		}
		// 邮箱不能为空
		if strings.TrimSpace(email) == "" {
			u.Redirect("/user/create", 302)
		}
		// 邮箱格式不合法
		// 密码加密
		// 保存
		// 回到用户详情
	}
	u.Data["roles"] = models.Roles()
	u.Layout = "base.html"
	u.TplName = "user/create.html"
}

// 用户查看
func (u *UserController) View() {

}

// 用户编辑
func (u *UserController) Update() {

}

// 删除用户
func (u *UserController) Delete() {

}

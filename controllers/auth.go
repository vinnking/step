// 用户认证
package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"step/models"
)

type AuthController struct {
	beego.Controller
}

// Login 管理员登陆
func (a *AuthController) Login() {
	if a.Ctx.Request.Method == "POST" {
		email := a.Input().Get("email")
		password := a.Input().Get("password")
		if strings.TrimSpace(email) == "" {
			a.Redirect("/auth/login", 302)
		}
		if strings.TrimSpace(password) == "" {
			a.Redirect("/auth/login", 302)
		}
		user, err := models.UserCheck(email, password)
		if user.Role != 1 || err != nil {
			a.Redirect("/auth/login", 302)
		}
		a.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), strconv.Itoa(int(user.Id)), 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		a.Redirect("/post", 302)
	}
	a.TplName = "login.html"
}

// Logout 管理员退出
func (a *AuthController) Logout() {
	a.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	a.Redirect("/auth/login", 302)
}

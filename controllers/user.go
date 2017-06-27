// 用户
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"step/models"
)

type UserController struct {
	beego.Controller
}

// 用户列表
func (u *UserController) Index() {
	u.Data["users"] = models.List()
	u.Data["status"] = models.Status()
	u.Data["roles"] = models.Roles()
	u.Layout = "base.html"
	u.TplName = "user/index.html"
}

// 添加用户
func (u *UserController) Create() {
	beego.ReadFromRequest(&u.Controller)
	if u.Ctx.Request.Method == "POST" {
		nickname := u.Input().Get("nickname")
		email := u.Input().Get("email")
		password := u.Input().Get("password")
		role, err := strconv.Atoi(u.Input().Get("role"))
		if err != nil {
			u.Redirect("/user/create", 302)
		}
		// 用户名成不能为空
		if strings.TrimSpace(nickname) == "" {
			u.Redirect("/user/create", 302)
		}
		// 邮箱不能为空
		if strings.TrimSpace(email) == "" {
			u.Redirect("/user/create", 302)
		}
		// 邮箱格式不合法
		// 密码不能为空
		if strings.TrimSpace(password) == "" {
			u.Redirect("/user/create", 302)
		}
		// 保存成功回到用户详情
		res, err := models.Save(&models.User{
			Nickname: nickname,
			Email:    email,
			Password: password,
			Role:     role,
		})
		if err != nil {
			u.Redirect("/user/create", 302)
		} else {
			u.Redirect("/user/"+strconv.FormatInt(res.Id, 10), 302)
		}
		fmt.Println("end save")
	}
	u.Data["roles"] = models.Roles()
	u.Data["user"] = models.User{}
	u.Data["isNewRecord"] = true
	u.Layout = "base.html"
	u.TplName = "user/create.html"
}

// 用户查看
func (u *UserController) View() {
	id, err := strconv.Atoi(u.Ctx.Input.Param(":id"))
	if err != nil {
		u.Redirect("/user", 302)
	}
	if id > 0 {
		user, err := models.Info(int64(id))
		if err != nil {
			u.Ctx.WriteString(fmt.Sprintf("用户%d信息为空", id))
		} else {
			u.Data["user"] = user
			u.Data["role"] = models.RoleDesc(user.Role)
			u.Data["status"] = models.StatusDesc(user.Status)
			u.Data["ctime"] = time.Unix(user.Ctime, 0).Format("2006-01-02 15:04:05")
			u.Data["utime"] = time.Unix(user.Utime, 0).Format("2006-01-02 15:04:05")
			u.Layout = "base.html"
			u.TplName = "user/view.html"
		}
	} else {
		u.Ctx.WriteString(fmt.Sprintf("用户%d不存在", id))
	}
}

// 用户编辑
func (u *UserController) Update() {
	id, err := strconv.Atoi(u.Ctx.Input.Param(":id"))
	if err != nil {
		u.Redirect("/user", 302)
	}
	if id > 0 {
		user, err := models.Info(int64(id))
		if err != nil {
			u.Ctx.WriteString(fmt.Sprintf("用户%d信息为空", id))
		} else {
			u.Data["user"] = user
			u.Data["isNewRecord"] = false
			u.Layout = "base.html"
			u.TplName = "/user/update.html"
		}
	} else {
		u.Ctx.WriteString(fmt.Sprintf("用户%d不存在", id))
	}
}

func (u *UserController) Stop() {

}

// 删除用户
func (u *UserController) Delete() {

}

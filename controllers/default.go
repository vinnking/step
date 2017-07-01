// 首页
package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"step/models"
)

type MainController struct {
	beego.Controller
}

// Get 首页
func (m *MainController) Get() {
	menuList := models.MenuList()

	var cateId int
	var err error
	cateId = 1
	id := m.Ctx.Input.Param(":id")
	if strings.TrimSpace(id) != "" {
		if cateId, err = strconv.Atoi(id); err != nil {
			cateId = 1
		}
		// 分类只获取到最大分类
		length := len(menuList)
		if cateId > length {
			cateId = length
		}
	}
	m.Data["menus"] = menuList
	m.Layout = "layout.html"
	m.TplName = "index.html"
}

// View 文章详情
func (m *MainController) View() {
	menuList := models.MenuList()
	m.Data["menus"] = menuList
	m.Layout = "layout.html"
	m.TplName = "view.html"
}

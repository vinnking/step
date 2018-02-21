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
	m.Data["quote"] = models.QuoteOne()
	m.Data["menus"] = menuList
	m.Data["cateId"] = cateId
	m.Data["posts"] = models.PostListFilter(cateId, 10)
	m.Data["recent"] = models.PostRecent()
	m.Data["labels"] = models.LabelList()
	m.Data["links"] = models.LinkList()
	m.Layout = "layout.html"
	m.TplName = "index.html"
}

// View 文章详情
func (m *MainController) View() {
	id, err := strconv.Atoi(m.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		m.Redirect("/", 302)
	}
	post, err := models.PostInfo(int64(id))
	if err != nil {
		m.Redirect("/", 302)
	}
	menuList := models.MenuList()
	m.Data["quote"] = models.QuoteOne()
	m.Data["menus"] = menuList
	m.Data["cateId"] = post.Type
	m.Data["recent"] = models.PostRecent()
	m.Data["labels"] = models.LabelList()
	m.Data["links"] = models.LinkList()
	m.Layout = "layout.html"
	m.TplName = post.Url
}

// Label 标签相关文章
func (m *MainController) Label() {
	menuList := models.MenuList()
	
	var labelId int
	var err error
	lId := m.Ctx.Input.Param(":id")
	if strings.TrimSpace(lId) != "" {
		if labelId, err = strconv.Atoi(lId); err != nil {
			labelId = 1
		}
	}
	
	m.Data["quote"] = models.QuoteOne()
	m.Data["menus"] = menuList
	m.Data["cateId"] = 0
	m.Data["posts"] = models.PostListLabel(labelId, 10)
	m.Data["recent"] = models.PostRecent()
	m.Data["labels"] = models.LabelList()
	m.Data["links"] = models.LinkList()
	m.Layout = "layout.html"
	m.TplName = "index.html"
}

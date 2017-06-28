// 首页
package controllers

import (
	"github.com/astaxie/beego"
	
	"step/models"
	"fmt"
	"strings"
	"strconv"
)

type MainController struct {
	beego.Controller
}

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
	fmt.Println(cateId)
	m.Data["menus"] = menuList
	m.Layout = "layout.html"
	m.TplName = "index.html"
}

// 友情链接
package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"step/models"
)

type LinkController struct {
	beego.Controller
}

// 检查用户是否登陆
func (l *LinkController) Prepare() {
	_, user := models.IsLogin(l.Ctx)
	l.Data["nickname"] = user.Nickname
}

// Index 友情链接列表
func (l *LinkController) Index() {
	l.Data["links"] = models.LinkList()
	l.Data["types"] = models.LinkTypeList()
	l.Data["status"] = models.LinkStatus()
	l.Layout = "base.html"
	l.TplName = "link/index.html"
}

// Create 添加友情链接
func (l *LinkController) Create() {
	if l.Ctx.Request.Method == "POST" {
		typeId, err := strconv.Atoi(strings.TrimSpace(l.Input().Get("type")))
		if err != nil {
			l.Redirect("/link/create", 302)
		}
		title := l.Input().Get("title")
		url := l.Input().Get("url")
		content := l.Input().Get("content")
		if typeId <= 0 {
			l.Redirect("/link/create", 302)
		}
		if strings.TrimSpace(title) == "" {
			l.Redirect("/link/create", 302)
		}
		if strings.TrimSpace(url) == "" {
			l.Redirect("/link/create", 302)
		}
		if strings.TrimSpace(content) == "" {
			l.Redirect("/link/create", 302)
		}
		var id int64
		if id, err = models.LinkSave(&models.Link{
			Type:    typeId,
			Title:   title,
			Url:     url,
			Content: content,
		}); err != nil {
			l.Redirect("/link/create", 302)
		}
		l.Redirect("/link/"+strconv.FormatInt(id, 10), 302)
	}

	l.Data["isNewRecord"] = true
	l.Data["link"] = models.Link{}
	l.Data["types"] = models.LinkTypeList()
	l.Layout = "base.html"
	l.TplName = "link/create.html"
}

// View 友情链接详情
func (l *LinkController) View() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/link", 302)
	}
	link, err := models.LinkInfo(int64(id))
	if err != nil {
		l.Redirect("/link ", 302)
	}
	l.Data["link"] = link
	l.Data["type"] = models.LinkTypeDesc(link.Type)
	l.Data["status"] = models.LinkStatusDesc(link.Status)
	l.Layout = "base.html"
	l.TplName = "link/view.html"
}

// Update 编辑友情链接
func (l *LinkController) Update() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/link", 302)
	}
	link, err := models.LinkInfo(int64(id))
	if err != nil {
		l.Redirect("/link ", 302)
	}

	if l.Ctx.Request.Method == "POST" {
		typeId, err := strconv.Atoi(strings.TrimSpace(l.Input().Get("type")))
		if err != nil {
			l.Redirect("/link/update/"+strconv.FormatInt(int64(id), 10), 302)
		}
		link.Type = typeId
		link.Title = strings.TrimSpace(l.Input().Get("title"))
		link.Url = strings.TrimSpace(l.Input().Get("url"))
		link.Content = strings.TrimSpace(l.Input().Get("content"))
		var newId int64
		if newId, err = models.LinkUpdate(&link); err != nil {
			l.Redirect("/link/update/"+strconv.FormatInt(newId, 10), 302)
		}
		l.Redirect("/link/"+strconv.FormatInt(newId, 10), 302)
	}

	l.Data["isNewRecord"] = false
	l.Data["link"] = link
	l.Data["types"] = models.LinkTypeList()
	l.Layout = "base.html"
	l.TplName = "link/update.html"
}

// Delete 删除友情链接
func (l *LinkController) Delete() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/link", 302)
	}
	link, err := models.LinkInfo(int64(id))
	if err != nil {
		l.Redirect("/link", 302)
	}
	link.Status = 2
	if _, err := models.LinkUpdate(&link); err != nil {
		fmt.Println(err)
	}
	l.Redirect("/link", 302)
}

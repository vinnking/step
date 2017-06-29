// 标签
package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"step/models"
)

type LabelController struct {
	beego.Controller
}

// Index 标签列表
func (l *LabelController) Index() {
	l.Data["labels"] = models.LabelList()
	l.Data["status"] = models.LinkStatus()
	l.Layout = "base.html"
	l.TplName = "label/index.html"
}

// Create 添加标签
func (l *LabelController) Create() {
	if l.Ctx.Request.Method == "POST" {
		name := l.Input().Get("name")
		content := l.Input().Get("content")
		if strings.TrimSpace(name) == "" {
			l.Redirect("/label/create", 302)
		}
		if strings.TrimSpace(content) == "" {
			l.Redirect("/label/create", 302)
		}
		var id int64
		var err error
		if id, err = models.LabelSave(&models.Label{
			Name:    name,
			Content: content,
		}); err != nil {
			l.Redirect("/label/create", 302)
		}
		l.Redirect("/label/"+strconv.FormatInt(id, 10), 302)
	}
	
	l.Data["isNewRecord"] = true
	l.Data["label"] = models.Label{}
	l.Layout = "base.html"
	l.TplName = "label/create.html"
}

// View 标签详情
func (l *LabelController) View() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/label", 302)
	}
	label, err := models.LabelInfo(int64(id))
	if err != nil {
		l.Redirect("/label ", 302)
	}
	l.Data["label"] = label
	l.Data["status"] = models.LabelStatusDesc(label.Status)
	l.Layout = "base.html"
	l.TplName = "label/view.html"
}

// Update 编辑标签
func (l *LabelController) Update() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/label", 302)
	}
	label, err := models.LabelInfo(int64(id))
	if err != nil {
		l.Redirect("/label ", 302)
	}
	
	if l.Ctx.Request.Method == "POST" {
		label.Name = strings.TrimSpace(l.Input().Get("name"))
		label.Content = strings.TrimSpace(l.Input().Get("content"))
		var newId int64
		if newId, err = models.LabelUpdate(&label); err != nil {
			l.Redirect("/label/update/"+strconv.FormatInt(newId, 10), 302)
		}
		l.Redirect("/label/"+strconv.FormatInt(newId, 10), 302)
	}
	
	l.Data["isNewRecord"] = false
	l.Data["label"] = label
	l.Layout = "base.html"
	l.TplName = "label/update.html"
}

// Delete 删除标签
func (l *LabelController) Delete() {
	id, err := strconv.Atoi(l.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		l.Redirect("/label", 302)
	}
	label, err := models.LabelInfo(int64(id))
	if err != nil {
		l.Redirect("/label", 302)
	}
	label.Status = 2
	if _, err := models.LabelUpdate(&label); err != nil {
		fmt.Println(err)
	}
	l.Redirect("/label", 302)
}

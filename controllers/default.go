package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Layout = "layout.html"
	m.TplName = "index.html"
}

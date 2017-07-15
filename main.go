package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"step/models"
	_ "step/routers"
	_ "step/templates"
)

func init() {
	orm.RegisterDataBase(
		"default",
		"mysql",
		beego.AppConfig.String("mysql.username")+":"+beego.AppConfig.String("mysql.password")+"@/"+beego.AppConfig.String("mysql.daname"),
	)
	orm.RegisterModel(
		new(models.User),
		new(models.Quote),
		new(models.Label),
		new(models.Link),
		new(models.Post),
		new(models.PostLabel),
	)
}

func main() {
	beego.Run()
}

package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"step/models"
	_ "step/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/step")
	orm.RegisterModel(
		new(models.User),
		new(models.Quote),
	)
}

func main() {
	beego.Run()
}

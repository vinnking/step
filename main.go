package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"step/models"
	_ "step/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1:3306)/step")
	orm.RegisterModel(
		new(models.User),
	)
}

func main() {
	beego.Run()
}

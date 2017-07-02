// 注册自定义模板函数
package templates

import (
	"time"

	"github.com/astaxie/beego"

	"step/models"
)

// UserName 模板渲染时根据用户标志获取用户名称(昵称)
func UserName(id int64) string {
	u, err := models.UserInfo(id)
	if err != nil {
		return "未知"
	}
	return u.Nickname
}

// ToDate 将int64格式的时间戳转化为日期字符串
func ToDate(t int64) string {
	return time.Unix(t, 0).Format("2006年01月02日")
}

func init() {
	beego.AddFuncMap("username", UserName)
	beego.AddFuncMap("todate", ToDate)
}

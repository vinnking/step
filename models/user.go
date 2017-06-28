// 用户模型
package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Nickname string
	Email    string
	Password string
	Salt     string
	Role     int
	Token    string
	Status   int
	Ctime    int64
	Utime    int64
}

// Roles 用户角色, 标识和描述
func Roles() map[int]string {
	return map[int]string{
		1: "管理员",
		2: "普通用户",
	}
}

// RoleDesc 根据角色标识获取角色描述
func RoleDesc(id int) string {
	if role, ok := Roles()[id]; ok {
		return role
	}
	return "未知"
}

// Status 用户状态
func Status() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// StatusDesc 用户状态描述
func StatusDesc(id int) string {
	if status, ok := Status()[id]; ok {
		return status
	}
	return "未知"
}

// Salt 生成一个盐, 用于校验密码复杂度
// 盐的长度不能超过120个字符
func Salt() string {
	return "cr42ew"
}

// Password 生成密码
func Password(password string, salt string) string {
	h := md5.New()
	h.Write([]byte(password + salt))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// Save 保存用户信息
// 返回用户信息和错误信息
func Save(u *User) (int64, error) {
	o := orm.NewOrm()
	u.Salt = Salt()
	u.Password = Password(u.Password, u.Salt)
	u.Status = 1
	u.Ctime = time.Now().Unix()
	u.Utime = time.Now().Unix()
	return o.Insert(u)
}

// Update 更新用户信息
func Update(u *User) (int64, error) {
	o := orm.NewOrm()
	// 如果密码不为空, 则修改密码
	if u.Password != "" {
		u.Salt = Salt()
		u.Password = Password(u.Password, u.Salt)
	}
	u.Utime = time.Now().Unix()
	if _, err := o.Update(u); err != nil {
		return u.Id, err
	}
	return u.Id, nil
}

// List 用户列表
func List() []*User {
	var user User
	var users []*User
	o := orm.NewOrm()
	o.QueryTable(user).RelatedSel().Filter("Status", 1).All(&users)
	return users
}

// Info 用户信息
func Info(id int64) (User, error) {
	var u User
	o := orm.NewOrm()
	err := o.QueryTable(u).RelatedSel().Filter("Id", id).One(&u)
	return u, err
}

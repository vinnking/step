package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int64
	Nickname string
	Email string
	Password string
	Role uint
	Token string
	Status uint
	Ctime int64
	Utime int64
}

// Roles 用户角色, 标识和描述
func Roles() map[int]string {
	return map[int]string{
		1 : "管理员",
		2 : "普通用户",
	}
}

// RoleDesc 根据角色标识获取角色描述
func RoleDesc(id int) string {
	if role, ok := Roles()[id]; ok {
		return role
	}
	return "未知"
}

// Save 保存用户信息
// 返回用户信息和错误信息
func (u *User) Save() (*User, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	if err != nil {
		return u, err
	}
	u.Id = id
	return u, nil
}

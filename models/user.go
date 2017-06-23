// 用户模型
package models

import (
	"crypto/md5"
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

// 密码强度
const strength = 16

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

// Salt 生成一个盐, 用于校验密码复杂度
// 盐的长度不能超过120个字符
func Salt() string {
	return "#-@,//a4H,>"
}

// Password 生成密码
// 密码强度至少为16次循环md5加密后加盐
func Password(password string, salt string, depth int) string {
	if depth < strength {
		depth = strength
	}
	md := md5.Sum([]byte(password))
	for i := 0; i < depth; i++ {
		md = md5.Sum([]byte(md))
	}
	return fmt.Sprintf("%s", md5.Sum([]byte(string(md)+salt)))
}

// Save 保存用户信息
// 返回用户信息和错误信息
func Save(u User) (User, error) {
	o := orm.NewOrm()
	u.Salt = Salt()
	u.Password = Password(u.Password, u.Salt, 16)
	u.Status = 1
	u.Ctime = time.Now().Unix()
	u.Utime = time.Now().Unix()
	id, err := o.Insert(u)
	if err != nil {
		return u, err
	}
	u.Id = id
	return u, nil
}

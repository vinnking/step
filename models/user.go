// 用户模型
package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"step/util"
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
func UserStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// StatusDesc 用户状态描述
func UserStatusDesc(id int) string {
	if status, ok := UserStatus()[id]; ok {
		return status
	}
	return "未知"
}

// Salt 生成一个盐, 用于校验密码复杂度
func Salt() string {
	s, err := util.Salt(6, 2, 2, 2)
	if err != nil {
		return "dMr$2H9"
	}
	return s
}

// Password 生成密码
func Password(password string, salt string) string {
	h := md5.New()
	h.Write([]byte(password + salt))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// Save 保存用户信息
// 返回用户信息和错误信息
func UserSave(u *User) (int64, error) {
	o := orm.NewOrm()
	u.Salt = Salt()
	u.Password = Password(u.Password, u.Salt)
	u.Status = 1
	u.Ctime = time.Now().Unix()
	u.Utime = time.Now().Unix()
	return o.Insert(u)
}

// Update 更新用户信息
func UserUpdate(u *User) (int64, error) {
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
func UserList() []*User {
	var user User
	var users []*User
	orm.NewOrm().QueryTable(user).RelatedSel().Filter("Status", 1).All(&users)
	return users
}

// Info 用户信息
func UserInfo(id int64) (User, error) {
	var u User
	err := orm.NewOrm().QueryTable(u).RelatedSel().Filter("Id", id).One(&u)
	return u, err
}

// UserCheck 检查用户是否存在
func UserCheck(email string, password string) (User, error) {
	var u User
	if err := orm.NewOrm().QueryTable(u).RelatedSel().Filter("Email", email).Filter("Status", 1).One(&u); err != nil {
		return u, err
	}
	// 核对密码是否正确
	if u.Password != Password(password, u.Salt) {
		return u, errors.New("密码错误")
	}
	return u, nil
}

// IsLogin 用户是否登陆, 无登陆直接跳转至登录页
func IsLogin(ctx *context.Context) (bool, User) {
	userId, ok := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	var user User
	if !ok {
		ctx.Redirect(302, "/auth/login")
	}
	newUserId, err := strconv.Atoi(userId)
	if err != nil {
		ctx.Redirect(302, "/auth/login")
	}
	user, err = UserInfo(int64(newUserId))
	if err != nil {
		ctx.Redirect(302, "/auth/login")
	}
	return ok, user
}

/*
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(120) NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(120) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(120) NOT NULL DEFAULT '#-@,//' COMMENT '盐',
  `role` tinyint(1) NOT NULL DEFAULT '2' COMMENT '角色1为管理员2为普通用户',
  `token` varchar(128) NOT NULL DEFAULT '' COMMENT '重置密码用',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1为可用0为不可用',
  `ctime` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `utime` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `role` (`role`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';
*/

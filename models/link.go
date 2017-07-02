// 友情链接
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Link struct {
	Id      int64
	Type    int
	Title   string
	Url     string
	Content string
	Status  int
	Ctime   int64
	Utime   int64
}

// LinkStatus 友情链接状态
func LinkStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// LinkStatusDesc 友情链接状态描述
func LinkStatusDesc(id int) string {
	if desc, ok := LinkStatus()[id]; ok {
		return desc
	}
	return "未知"
}

// LinkTypeList 类型
func LinkTypeList() map[int]string {
	return map[int]string{
		1: "官网",
		2: "手册",
		3: "其它",
	}
}

// LinkTypeDesc 类型描述
func LinkTypeDesc(id int) string {
	if desc, ok := LinkTypeList()[id]; ok {
		return desc
	}
	return "未知"
}

// LinkSave 添加友情链接
func LinkSave(l *Link) (int64, error) {
	l.Status = 1
	l.Ctime = time.Now().Unix()
	l.Utime = time.Now().Unix()
	return orm.NewOrm().Insert(l)
}

// LinkUpdate 更新友情链接
func LinkUpdate(l *Link) (int64, error) {
	l.Utime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(l); err != nil {
		return l.Id, err
	}
	return l.Id, nil
}

// LinkList 友情链接列表
func LinkList() []*Link {
	var link Link
	var links []*Link
	orm.NewOrm().QueryTable(link).RelatedSel().Filter("Status", 1).All(&links)
	return links
}

// LinkInfo 友情链接详情
func LinkInfo(id int64) (Link, error) {
	var l Link
	err := orm.NewOrm().QueryTable(l).RelatedSel().Filter("Id", id).One(&l)
	return l, err
}

/*
CREATE TABLE `link` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型1官网2手册3其它',
  `title` varchar(120) NOT NULL COMMENT '标题',
  `url` varchar(255) NOT NULL COMMENT '地址',
  `content` varchar(255) NOT NULL COMMENT '描述',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1可用2不可用',
  `ctime` int(11) DEFAULT '0' COMMENT '创建时间',
  `utime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `type`(`type`),
  KEY `title`(`title`),
  KEY `status`(`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='友情链接';
*/

// 引用
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Quote struct {
	Id      int64
	Author  string
	Content string
	Extra   string
	Status  int
	Ctime   int64
	Utime   int64
}

// QuoteStatus 引用状态
func QuoteStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// QuoteStatusDesc 引用状态描述
func QuoteStatusDesc(id int) string {
	if desc, ok := QuoteStatus()[id]; ok {
		return desc
	}
	return "未知"
}

// QuoteSave 添加引用
func QuoteSave(q *Quote) (int64, error) {
	q.Status = 1
	q.Ctime = time.Now().Unix()
	q.Utime = time.Now().Unix()
	return orm.NewOrm().Insert(q)
}

// QuoteUpdate 更新引用
func QuoteUpdate(q *Quote) (int64, error) {
	q.Utime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(q); err != nil {
		return q.Id, err
	}
	return q.Id, nil
}

// QuoteList 引用列表
func QuoteList() []*Quote {
	var quote Quote
	var quotes []*Quote
	orm.NewOrm().QueryTable(quote).RelatedSel().Filter("Status", 1).All(&quotes)
	return quotes
}

// QuoteInfo 引用详情
func QuoteInfo(id int64) (Quote, error) {
	var q Quote
	err := orm.NewOrm().QueryTable(q).RelatedSel().Filter("Id", id).One(&q)
	return q, err
}

// QuoteOne 获取最新的一条引用
func QuoteOne() Quote {
	var q Quote
	if err := orm.NewOrm().QueryTable(q).RelatedSel().Filter("Status", 1).OrderBy("-Id").One(&q); err != nil {
		// 需要添加引用
	}
	return q
}

/*
CREATE TABLE `quote` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `author` varchar(64) NOT NULL COMMENT '作者',
  `content` varchar(255) NOT NULL COMMENT '引用内容',
  `extra` varchar(255) NOT NULL COMMENT '引用描述, 解说',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1可用2不可用',
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `utime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status`(`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
*/

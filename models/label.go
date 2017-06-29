// 标签管理
package models

import (
	"github.com/astaxie/beego/orm"
)

type Label struct {
	Id      int64
	Name    string
	Content string
	Status  int
}

// LabelStatus 标签状态
func LabelStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// LabelStatusDesc 标签状态描述
func LabelStatusDesc(id int) string {
	if desc, ok := LabelStatus()[id]; ok {
		return desc
	}
	return "未知"
}

// LabelSave 添加标签
func LabelSave(l *Label) (int64, error) {
	o := orm.NewOrm()
	l.Status = 1
	return o.Insert(l)
}

// LabelUpdate 更新标签
func LabelUpdate(l * Label) (int64, error) {
	o := orm.NewOrm()
	if _, err := o.Update(l); err != nil {
		return l.Id, err
	}
	return l.Id, nil
}

// LabelList 标签列表
func LabelList() []*Label {
	var label Label
	var labels []*Label
	o := orm.NewOrm()
	o.QueryTable(label).RelatedSel().Filter("Status", 1).All(&labels)
	return labels
}

// LabelInfo 标签详情
func LabelInfo(id int64) (Label, error) {
	var l Label
	o := orm.NewOrm()
	err := o.QueryTable(l).RelatedSel().Filter("Id", id).One(&l)
	return l, err
}

/**
CREATE TABLE `label` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
 `name` varchar(120) NOT NULL,
 `content` varchar(255) NOT NULL DEFAULT '' COMMENT '标签描述',
 `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态1可用2不可用',
 KEY `name`(`name`),
 KEY `status`(`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '标签管理';
*/

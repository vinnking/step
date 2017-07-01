// 文章标签关联
package models

import "github.com/astaxie/beego/orm"

type PostLabel struct {
	Id      int64
	LabelId int64
	PostId  int64
	UserId  int64
}

// PostLabelList 根据用户和文章查询关联的标签
func PostLabelList(userId int64, postId int64) []*PostLabel {
	var postLabel PostLabel
	var postLabels []*PostLabel
	o := orm.NewOrm()
	o.QueryTable(postLabel).RelatedSel().Filter("UserId", userId).Filter("PostId", postId).All(&postLabels)
	return postLabels
}

// PostLabelSave 文章标签关联保存
func PostLabelSave(pl *PostLabel) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(pl)
}

// PostLabelDelete 文章标签关联删除
func PostLabelDelete(userId int64, postId int64) (int64, error) {
	var pl PostLabel
	o := orm.NewOrm()
	return o.QueryTable(pl).RelatedSel().Filter("UserId", userId).Filter("PostId", postId).Delete()
}

/*
CREATE TABLE `post_label` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
 `label_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '标签记录',
 `post_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '文章记录',
 `user_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户记录',
 UNIQUE `label_post_user`(`label_id`, `post_id`, `user_id`),
 KEY `label_id`(`label_id`),
 KEY `user_id`(`user_id`),
 KEY `post_id`(`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '文章标签关联';
*/

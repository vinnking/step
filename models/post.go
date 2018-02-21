// 文章
package models

import (
	"fmt"
	"time"
	
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id        int64
	Title     string
	Type      int
	Url       string
	ImageUrl  string
	Summary   string
	UserId    int64
	ReadCount int32
	Remark    string
	Status    int
	Ctime     int64
	Utime     int64
}

// PostStatus 文章状态
func PostStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// PostStatusDesc 文章状态描述
func PostStatusDesc(id int) string {
	if desc, ok := LinkStatus()[id]; ok {
		return desc
	}
	return "未知"
}

// PostSave 文章保存
func PostSave(p *Post) (int64, error) {
	o := orm.NewOrm()
	p.Status = 1
	p.Ctime = time.Now().Unix()
	p.Utime = time.Now().Unix()
	return o.Insert(p)
}

// PostUpdate 文章更新
func PostUpdate(p *Post) (int64, error) {
	p.Utime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(p); err != nil {
		return p.Id, err
	}
	return p.Id, nil
}

// PostList 文章列表
func PostList() []*Post {
	var post Post
	var posts []*Post
	orm.NewOrm().QueryTable(post).RelatedSel().Filter("Status", 1).OrderBy("-Id").All(&posts)
	return posts
}

// PostListFilter 根据分类和条数获取文章列表
func PostListFilter(cateId int, limit int) []*Post {
	var post Post
	var posts []*Post
	orm.NewOrm().QueryTable(post).RelatedSel().Filter("Status", 1).Filter("Type", cateId).OrderBy("-Id").Limit(limit).All(&posts)
	return posts
}

// PostListLabel 根据标签搜索文章
func PostListLabel(labelId int, limit int) []*Post {
	var postIds []int64
	// 该标签下的所有文章id
	for _, label := range PostLabels(int64(labelId)) {
		postIds = append(postIds, label.PostId)
	}
	var post Post
	var posts []*Post
	if len(postIds) > 0 {
		orm.NewOrm().QueryTable(post).RelatedSel().Filter("Id__in", postIds).Filter("Status", 1).OrderBy("-Id").Limit(limit).All(&posts)
	}
	return posts
}

// PostInfo 文章详情
func PostInfo(id int64) (Post, error) {
	var post Post
	err := orm.NewOrm().QueryTable(post).RelatedSel().Filter("Id", id).One(&post)
	return post, err
}

// PostLabelMap 文章关联的标签关联字典
func PostLabelMap(userId int64, postId int64) map[int64]string {
	desc := make(map[int64]string)
	labels := make(map[int64]string)
	for _, label := range LabelList() {
		labels[label.Id] = label.Name
	}
	for _, postLabel := range PostLabelList(userId, postId) {
		if d, ok := labels[postLabel.LabelId]; ok {
			desc[postLabel.LabelId] = d
		}
	}
	return desc
}

// PostLabelDesc 文章关联的标签描述
func PostLabelDesc(userId int64, postId int64) string {
	var desc string
	for id, label := range PostLabelMap(userId, postId) {
		desc += fmt.Sprintf("(%d - %s) ", id, label)
	}
	return desc
}

// PostRecent 最新的三篇文章
func PostRecent() []*Post {
	var post Post
	var posts []*Post
	orm.NewOrm().QueryTable(post).RelatedSel().Filter("Status", 1).OrderBy("-Id").Limit(3).All(&posts)
	return posts
}

/*
CREATE TABLE `post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(120) NOT NULL COMMENT '标题',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '文件路径',
  `image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片路径',
  `summary` varchar(255) NOT NULL COMMENT '摘要',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '作者',
  `read_count` int(11) NOT NULL DEFAULT '0' COMMENT '阅读次数',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1可用2不可用',
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `utime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  KEY `title`(`title`),
  KEY `type`(`type`),
  KEY `user_id`(`user_id`),
  KEY `status`(`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章管理';
*/

// 文章标签关联
package models

/**
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

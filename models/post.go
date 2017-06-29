package models

type Post struct {

}

/**
CREATE TABLE `post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(120) NOT NULL,
  `type` tinyint(1) NOT NULL DEFAULT '1',
  `content` mediumtext NOT NULL,
  `summary` varchar(255) NOT NULL COMMENT '摘要',
  `user_id` int(11) NOT NULL DEFAULT '0',
  `read_count` int(11) NOT NULL DEFAULT '0',
  `remark` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `ctime` int(11) NOT NULL DEFAULT '0',
  `utime` int(11) NOT NULL DEFAULT '0',
  KEY `title`(`title`),
  KEY `status`(`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章管理';
 */
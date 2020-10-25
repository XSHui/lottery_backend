package model

type User struct {
	Id       string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	PhoneNumber uint64 `json:"phone_number" xorm:"not null unique default 0 INT(11)"`
	// TODO: State
	CreateTime  int    `json:"create_time" xorm:"INT(10)"`
	ModifyTime  int    `json:"modify_time" xorm:"INT(10)"`
	DeleteTime  int    `json:"delete_time" xorm:"INT(10)"`
}

// 35 CREATE TABLE IF NOT EXISTS `lottery_users` (
// 36  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
// 37  `phone` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '手机号',
// 38  `draw_right` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可以抽奖 0否 1是',
// 39  `article` varchar(500) NOT NULL DEFAULT '0' COMMENT '文章',
// 40  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 41  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
// 42  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
// 43  PRIMARY KEY (`id`),
// 44  UNIQUE KEY `uk_phone` (`phone`),
// 45  KEY `update_time` (`update_time`)
// 46 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户报名信息';

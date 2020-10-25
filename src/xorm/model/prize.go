package model

type Prize struct {
	Id string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	// TODO: ActivityId
	Award int `json:"award" xorm:"INT(10)"`
	Total int `json:"total" xorm:"INT(10)"`
	Left  int `json:"left" xorm:"INT(10)"`
	// TODO: State
	CreateTime int `json:"create_time" xorm:"INT(10)"`
	ModifyTime int `json:"modify_time" xorm:"INT(10)"`
	DeleteTime int `json:"delete_time" xorm:"INT(10)"`
}

//4 CREATE TABLE IF NOT EXISTS `lottery_prizes` (
//5  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
//6  `prize` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖品类型 0-贴纸 1-电话卡 2-手机',
//7  `total` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品总数',
//8  `stock` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品剩余数量',
//9  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//10  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
//11  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
//12  PRIMARY KEY (`id`),
//13  UNIQUE KEY `uk_prize` (`prize`),
//14  KEY `update_time` (`update_time`)
//15 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='奖品列表';

package model

type Record struct {
	Id     string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	UserId string `json:"user_id" xorm:"not null unique(permission) default '' VARCHAR(128)"`
	// TODO: ActivityId
	PrizeId string `json:"user_id" xorm:"not null unique(permission) default '' VARCHAR(128)"`
	// TODO: State
	CreateTime int `json:"create_time" xorm:"not null unique(permission) default 0 INT(10)"`
	ModifyTime int `json:"modify_time" xorm:"INT(10)"`
	DeleteTime  int    `json:"delete_time" xorm:"INT(10)"`
}

//21 CREATE TABLE IF NOT EXISTS `lottery_records` (
//22  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
//23  `phone` varchar(11) NOT NULL DEFAULT '0' COMMENT '手机号',
//24  `prize_id` int(11) NOT NULL DEFAULT '0' COMMENT 'prize',
//25  `draw_date` date NOT NULL DEFAULT '1970-01-01' COMMENT '抽奖日期',
//26  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//27  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
//28  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
//29  PRIMARY KEY (`id`),
//30  KEY `phone_prize_id_draw_date` (`phone`, `prize_id`, `draw_date`),
//31  KEY `update_time` (`update_time`)
//32 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抽奖记录';

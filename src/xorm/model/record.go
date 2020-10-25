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



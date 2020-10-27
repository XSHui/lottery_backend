package model

type Prize struct {
	Id string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	// TODO: ActivityId
	Award     uint64  `json:"award" xorm:"INT(10)"`
	Name      string  `json:"prize name" xorm:"VARCHAR(256)"`
	Odds      float32 `json:"odds" xorm:"FLOAT"`
	Total     int     `json:"total" xorm:"INT(10)"`
	Left      int     `json:"left" xorm:"INT(10)"`
	Unlimited int     `json:"unlimited" xorm:"default 0 TINYINT(1)"`
	// TODO: State
	CreateTime int `json:"create_time" xorm:"INT(10)"`
	ModifyTime int `json:"modify_time" xorm:"INT(10)"`
	DeleteTime int `json:"delete_time" xorm:"INT(10)"`
}

package model

type Prize struct {
	Id string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	// TODO: ActivityId
	Name      string  `json:"prize name" xorm:"VARCHAR(256)"`
	OddsStart float32 `json:"odds_start" xorm:"FLOAT"`
	OddsEnd   float32 `json:"odds_end" xorm:"FLOAT"`
	Total     int     `json:"total" xorm:"INT(10)"`
	Left      int     `json:"left" xorm:"INT(10)"` // TODO: left is a key in mysql
	Unlimited int     `json:"unlimited" xorm:"default 0 TINYINT(1)"`
	DayLimit  int     `json:"day_limit" xorm:"INT(10)"`
	UserLimit int     `json:"user_limit" xorm:"INT(10)"`
	// TODO: State
	CreateTime int `json:"create_time" xorm:"INT(10)"`
	ModifyTime int `json:"modify_time" xorm:"INT(10)"`
	DeleteTime int `json:"delete_time" xorm:"INT(10)"`
}

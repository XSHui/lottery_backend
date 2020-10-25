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



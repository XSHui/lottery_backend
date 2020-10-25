package model

type User struct {
	Id       string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	PhoneNumber uint64 `json:"phone_number" xorm:"not null unique default 0 INT(11)"`
	// TODO: State
	CreateTime  int    `json:"create_time" xorm:"INT(10)"`
	ModifyTime  int    `json:"modify_time" xorm:"INT(10)"`
	DeleteTime  int    `json:"delete_time" xorm:"INT(10)"`
}


package model

type Permission struct {
	UserId string `json:"user_id" xorm:"not null unique default '' VARCHAR(128)"`
	// TODO: ActivityId
	//PhoneNumber uint64 `json:"phone_number" xorm:"not null unique default 0 INT(11)"`
	Permitted int `json:"permitted" xorm:"default 0 TINYINT(1)"`
}

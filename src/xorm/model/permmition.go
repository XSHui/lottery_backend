package model

type Permission struct {
	UserId       string `json:"user_id" xorm:"not null unique default '' VARCHAR(128)"`
	// TODO: ActivityId
	Permitted    int `json:"permitted" xorm:"default 0 TINYINT(1)"`
}


package model

type Article struct {
	Id     string `json:"id" xorm:"not null pk default '' VARCHAR(128)"`
	UserId string `json:"user_id" xorm:"not null default '' VARCHAR(128)"`
	// TODO: ActivityId
	Text   string `json:"article" xorm:"not null default '' VARCHAR(1024)"`
}

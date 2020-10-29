package xorm

import (
	"errors"
	"time"

	"lottery_backend/src/xorm/model"
)

// InsertRecord: insert lottery record to db
func InsertRecord(record *model.Record) error {
	if record == nil || record.UserId == "" {
		return errors.New("insert record error")
	}

	db := GetDB()
	_, err := db.Insert(record)
	return err
}

// ListRecord: list all prize record
func ListRecord(offset, limit int) ([]model.Record, error) {
	db := GetDB()
	data := make([]model.Record, 0)
	err := db.Limit(limit, offset).Find(&data)
	return data, err
}

// DaylotteryCount: user record count in one day
func DaylotteryCount(userId string) (int, error) {
	if userId == "" {
		return 0, errors.New("DaylotteryCount input invalid")
	}

	// 00:00
	t := time.Now()
	zeroTm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()

	db := GetDB()
	record := new(model.Record)
	total, err := db.Where("user_id = ? and create_time > ?", userId, zeroTm).Count(record)
	return int(total), err
}

// PrizeDayCount: day give out count of prize
func PrizeDayCount(prizeId string) (int, error) {
	if prizeId == "" {
		return 0, errors.New("PrizeDayCount input invalid")
	}
	// 00:00
	t := time.Now()
	zeroTm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()

	db := GetDB()
	record := new(model.Record)
	total, err := db.Where("prize_id = ? and create_time > ?", prizeId, zeroTm).Count(record)
	return int(total), err
}

// PrizeUserCount: user give out of prize
func PrizeUserCount(prizeId, userId string) (int, error) {
	if prizeId == "" || userId == "" {
		return 0, errors.New("PrizeUserCount input invalid")
	}

	db := GetDB()
	record := new(model.Record)
	total, err := db.Where("prize_id = ? and user_id = ?", prizeId, userId).Count(record)
	return int(total), err
}

// SubOneDayForRecord
func SubOneDayForRecord() error {
	db := GetDB()
	_, err := db.Exec("update record set create_time = create_time - 86400 and modify_time = modify_time - 86400")
	return err
}

package xorm

import (
	"errors"
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
func ListRecord(offset, limit int)([]model.Record, error) {
	db := GetDB()
	data := make([]model.Record, 0)
	err := db.Limit(limit, offset).Find(&data)
	return data, err
}

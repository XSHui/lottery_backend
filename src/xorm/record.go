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

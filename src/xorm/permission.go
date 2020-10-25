package xorm

import (
	"errors"
	"lottery_backend/src/xorm/model"
)

// InserPermission: prize right
func InserPermission(permission *model.Permission) error {
	if permission == nil || permission.UserId == "" {
		return errors.New("insert permission error")
	}

	db := GetDB()
	_, err := db.Insert(permission)
	return err
}
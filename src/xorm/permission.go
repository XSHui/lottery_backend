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

// Permitted: permitted is true or false
func Permitted(userId string) (bool, error) {
	db := GetDB()
	permission := model.Permission{UserId: userId}
	exist, err := db.Where("user_id = ?", permission.UserId).Get(&permission)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return permission.Permitted == 1, nil
}

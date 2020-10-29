package xorm

import (
	"errors"
	"lottery_backend/src/xorm/model"
)

// InsertUser: insert user to db
func InsertUser(user *model.User) error {
	if user == nil || user.PhoneNumber == 0 || user.Id == ""{
		return errors.New("insert user error")
	}

	db := GetDB()
	_, err := db.Insert(user)
	return err
}

// GetUserInfoByPhoneNum: get user info by phone number
func GetUserInfoByPhoneNum(phone uint64) (*model.User, error) {
	if phone == 0 {
		return nil, errors.New("invalid phone number")
	}

	db := GetDB()
	data := make([]model.User, 0)
	err := db.Where("phone_number=?", phone).Find(&data)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, errors.New("phone number not the only")
	}
	return &data[0], nil
}

// ListUser
func ListUser(offset, limit int)([]model.User, error) {
	db := GetDB()
	data := make([]model.User, 0)
	err := db.Limit(limit, offset).Find(&data)
	return data, err
}







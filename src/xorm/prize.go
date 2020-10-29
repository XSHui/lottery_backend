package xorm

import (
	"lottery_backend/src/xorm/model"
)

// GetPrize: get all peize
// TODO: GetPrizeByActivityId
func GetPrize() ([]model.Prize, error) {
	db := GetDB()
	data := make([]model.Prize, 0)
	err := db.Find(&data)
	return data, err
}

//// GetMaxCountPrize: get maxcount prize
//// TODO: GetPrizeByActivityId
//func GetMaxCountPrize() (model.Prize, error) {
//	db := GetDB()
//	data := make([]model.Prize, 0)
//	err := db.Desc("odds").Limit(1, 0).Find(&data)
//	if err != nil {
//		return model.Prize{}, err
//	}
//	if len(data) != 1 {
//		return model.Prize{}, errors.New("no prize get from db")
//	}
//	return data[0], nil
//}

// UpdatePrizeLeft:
// TODO: UpdatePrizeLeft(prizeId, ActivityId)
func UpdatePrizeLeft(prizeId string) error {
	db := GetDB()
	_, err := db.Exec("update prize set `left` = `left` -1 where `left` > 0 and id = ?", prizeId)
	return err
}

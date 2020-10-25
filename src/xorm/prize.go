package xorm

import "lottery_backend/src/xorm/model"

// GetPrize: get all peize
// TODO: GetPrizeByActivityId
func GetPrize()([]model.Prize, error) {
	db := GetDB()
	data := make([]model.Prize, 0)
	err := db.Find(&data)
	return data, err
}

// UpdatePrizeLeft:
// TODO: UpdatePrizeLeft(award, ActivityId)
func UpdatePrizeLeft(award int) error {
	db := GetDB()
	_, err := db.Exec("update prize set left = left -1 where left > 0 and award = ?", award)
	return err
}

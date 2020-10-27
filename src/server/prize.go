package api

import (
	cfg "lottery_backend/src/config"
	"lottery_backend/src/xorm"
	xmodel "lottery_backend/src/xorm/model"
)

func lotteryPrize(lotteryCode int) (xmodel.Prize, error) {
	// TODO: to be more robust
	prize, err := xorm.GetMaxCountPrize()
	if err != nil {
		return xmodel.Prize{}, err
	}

	prizes, err := xorm.GetPrize()
	if err != nil {
		return xmodel.Prize{}, err
	}

	for _, p := range prizes {
		if lotteryCode >= int(cfg.BASE_LOTTERY_CODE*p.Odds) && lotteryCode <= int(cfg.BASE_LOTTERY_CODE*p.Odds) {
			return p, nil
		}
	}

	return prize, nil
}

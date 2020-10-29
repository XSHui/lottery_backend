package api

import (
	"errors"

	cfg "lottery_backend/src/config"
	"lottery_backend/src/xorm"
	xmodel "lottery_backend/src/xorm/model"
)

// lotteryPrize: lottery
func lotteryPrize(lotteryCode int) (xmodel.Prize, error) {
	prizes, err := xorm.GetPrize()
	if err != nil {
		return xmodel.Prize{}, err
	}

	if len(prizes) == 0 {
		return xmodel.Prize{}, errors.New("no prize got from db")
	}

	// probability
	for _, p := range prizes {
		if lotteryCode >= int(cfg.BASE_LOTTERY_CODE*p.OddsStart) && lotteryCode <= int(cfg.BASE_LOTTERY_CODE*p.OddsEnd) {
			return p, nil
		}
	}

	return xmodel.Prize{}, nil
}

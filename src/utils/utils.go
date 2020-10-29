package utils

import (
	"math/rand"
	"time"

	"github.com/satori/go.uuid"

	"lottery_backend/src/config"
)

// NewId: new id, uuid
func NewId() string {
	return uuid.NewV4().String()
}

// NowTimestamp: time stamp
func NowTimestamp() int {
	return int(time.Now().Unix())
}

// NewRedisLockId: new redis lock id, uuid
func NewRedisLockId() string {
	return uuid.NewV4().String()
}

// NewLotteryCode: new lottery code
func NewLotteryCode() int {
	seed := time.Now().UnixNano() // seed
	r := rand.New(rand.NewSource(seed))
	return r.Intn(config.BASE_LOTTERY_CODE)
}

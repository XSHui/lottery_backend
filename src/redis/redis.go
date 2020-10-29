package dao

import (
	"sync"

	"github.com/gomodule/redigo/redis"

	"lottery_backend/src/xlog"
)

type RedisManager struct {
	RedisPool *redis.Pool
}

var once sync.Once
var redisManager *RedisManager = nil

func GetRedisInstance() *RedisManager {
	once.Do(func() {
		redisManager = new(RedisManager)
		redisManager.RedisPool = new(redis.Pool)
	})
	return redisManager
}

const (
	SET_IF_NOT_EXIST  = "NX" //SETNX key value
	SET_EXPIRE_TIME   = "PX" // millisecond
	PX_TIME_OUT       = 100  // 100 ms
	LOCK_OK           = "OK" // get lock success
	DELLOCK_OK        = 1    // unlock success
	DELLOCK_NOT_EXIST = 0    // lock does not exits when unlock
)

func (rm *RedisManager) SetLock(key, requestId string) bool {
	conn := rm.RedisPool.Get()
	defer conn.Close()
	msg, err := redis.String(
		conn.Do("SET", key, requestId, SET_IF_NOT_EXIST, SET_EXPIRE_TIME, PX_TIME_OUT),
	)
	if err != redis.ErrNil && err != nil {
		xlog.ErrorSimple("SetLock ERROR", xlog.Fields{
			"key":       key,
			"requestId": requestId,
			"err":       err,
		})
	}
	if msg == LOCK_OK {
		return true
	}
	return false
}

func (rm *RedisManager) UnLock(key, requestId string) bool {
	conn := rm.RedisPool.Get()
	defer conn.Close()
	if rm.GetLock(conn, key) == requestId {
		msg, err := redis.Int64(conn.Do("DEL", key))
		if err != redis.ErrNil && err != nil {
			xlog.ErrorSimple("DELLOCK ERROR", xlog.Fields{
				"key":       key,
				"requestId": requestId,
				"err":       err,
			})
		}
		// time out
		if msg == DELLOCK_OK || msg == DELLOCK_NOT_EXIST {
			return true
		}
		return false
	}
	return false
}

func (rm *RedisManager) GetLock(conn redis.Conn, key string) string {
	msg, err := redis.String(conn.Do("GET", key))
	if err != redis.ErrNil && err != nil {
		xlog.ErrorSimple("GetLock ERROR", xlog.Fields{
			"key": key,
			"err": err,
		})
	}
	return msg
}
